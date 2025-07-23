package main

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/gin-gonic/gin"
	"github.com/morehao/golib/protocol"
	"github.com/morehao/golib/protocol/gresty"
	"github.com/morehao/golib/storages/dbes"
)

func singleEmbedding(ctx *gin.Context, content string) (embedding []float32, err error) {
	httpClientCfg := &protocol.HttpClientConfig{
		Module: "esvector",
		Host:   EmbeddingHost,
	}
	restyClient := gresty.NewClient(httpClientCfg)
	embeddingReq := map[string]any{
		"model": "Qwen3-Embedding-0.6B",
		"input": content,
	}
	var embeddingRes EmbeddingResponse
	request, newRequestErr := restyClient.NewRequestWithResult(ctx, &embeddingRes)
	if newRequestErr != nil {
		return nil, newRequestErr
	}
	_, err = request.SetBody(embeddingReq).Post("/v1/embeddings")
	if err != nil {
		return nil, err
	}
	if len(embeddingRes.Data) == 0 {
		return nil, fmt.Errorf("embeddingRes.Data is empty, content: %s", content)
	}
	return embeddingRes.Data[0].Embedding, nil

}

// executeSearch 通用的搜索执行函数
func executeSearch(ctx *gin.Context, queryBuilder *dbes.Builder) (*SearchResponse, error) {
	if queryBuilder == nil {
		return nil, fmt.Errorf("query builder is nil")
	}

	searchBody, buildErr := queryBuilder.BuildReader()
	if buildErr != nil {
		return nil, fmt.Errorf("failed to build search query: %w", buildErr)
	}

	searchRes, searchErr := ESClient.Search(
		ESClient.Search.WithContext(ctx),
		ESClient.Search.WithIndex(ESIndexName),
		ESClient.Search.WithBody(searchBody))

	if searchErr != nil {
		return nil, fmt.Errorf("search request failed: %w", searchErr)
	}

	// 确保关闭响应体
	if searchRes != nil && searchRes.Body != nil {
		defer searchRes.Body.Close()
	}

	if searchRes.StatusCode != 200 {
		return nil, fmt.Errorf("search failed with status: %s", searchRes.Status())
	}

	var res ESSearchResponse
	if err := json.NewDecoder(searchRes.Body).Decode(&res); err != nil {
		return nil, fmt.Errorf("failed to decode search response: %w", err)
	}

	return &SearchResponse{
		Total:    res.Hits.Total.Value,
		MaxScore: res.Hits.MaxScore,
		List:     res.Hits.Hits,
	}, nil
}

// textSearch 文本搜索
func textSearch(ctx *gin.Context, searchValue string) (*SearchResponse, error) {

	cfg := DefaultSearchConfig

	queryBuilder := dbes.NewBuilder().
		SetSource([]string{"doc_id", "content", "category"}).
		SetSize(cfg.K).
		SetQuery(dbes.BuildMap("match", dbes.BuildMap("content", searchValue)))

	return executeSearch(ctx, queryBuilder)
}

// vectorSearch 向量搜索
func vectorSearch(ctx *gin.Context, searchValue string) (*SearchResponse, error) {

	cfg := DefaultSearchConfig

	embedding, embeddingErr := singleEmbedding(ctx, searchValue)
	if embeddingErr != nil {
		return nil, fmt.Errorf("failed to generate embedding: %w", embeddingErr)
	}

	queryBuilder := dbes.NewBuilder().
		SetSource([]string{"doc_id", "content", "category"}).
		Set("knn", dbes.BuildMap(
			"field", "embedding",
			"query_vector", embedding,
			"k", cfg.K,
			"num_candidates", cfg.NumCandidates))

	return executeSearch(ctx, queryBuilder)
}

// textVectorSearch 文本+向量搜索
func textVectorSearch(ctx *gin.Context, searchValue string) (*SearchResponse, error) {

	cfg := DefaultSearchConfig

	embedding, embeddingErr := singleEmbedding(ctx, searchValue)
	if embeddingErr != nil {
		return nil, fmt.Errorf("failed to generate embedding: %w", embeddingErr)
	}

	// 最简洁的混合搜索：knn和query同级，ES自动合并
	queryBuilder := dbes.NewBuilder().
		SetSource([]string{"doc_id", "content", "category"}).
		SetSize(cfg.K).
		Set("knn", dbes.BuildMap(
			"field", "embedding",
			"query_vector", embedding,
			"k", cfg.K,
			"num_candidates", cfg.NumCandidates)).
		SetQuery(dbes.BuildMap("match", dbes.BuildMap(
			"content", dbes.BuildMap(
				"query", searchValue))))

	return executeSearch(ctx, queryBuilder)
}

func hybridSearchByScriptScore(ctx *gin.Context, searchValue string) (*SearchResponse, error) {

	cfg := DefaultSearchConfig

	embedding, embeddingErr := singleEmbedding(ctx, searchValue)
	if embeddingErr != nil {
		return nil, fmt.Errorf("failed to generate embedding: %w", embeddingErr)
	}

	textQuery := dbes.BuildMap("match", dbes.BuildMap("content", searchValue))

	scriptScoreQuery := dbes.BuildMap("script_score",
		dbes.BuildMap("query", textQuery,
			"script",
			dbes.BuildMap("source", "cosineSimilarity(params.query_vector, 'embedding') + 1.0", "params", dbes.BuildMap("query_vector", embedding))))

	queryBuilder := dbes.NewBuilder().
		SetSource([]string{"doc_id", "content", "category"}).
		SetSize(cfg.K).
		SetQuery(scriptScoreQuery)

	return executeSearch(ctx, queryBuilder)
}

func hybridSearchByMemory(ctx *gin.Context, searchValue string) (*SearchResponse, error) {

	textSearchRes, textSearchErr := textSearch(ctx, searchValue)
	if textSearchErr != nil {
		return nil, fmt.Errorf("failed to text search: %w", textSearchErr)
	}

	vectorSearchRes, vectorSearchErr := vectorSearch(ctx, searchValue)
	if vectorSearchErr != nil {
		return nil, fmt.Errorf("failed to vector search: %w", vectorSearchErr)
	}

	// 融合结果
	const (
		textWeight   = 0.3
		vectorWeight = 0.7
	)

	dataMap := make(map[string]ESSearchHitsItem, len(textSearchRes.List)+len(vectorSearchRes.List))

	// 处理文本搜索结果
	for _, item := range textSearchRes.List {
		item.Score *= textWeight
		dataMap[item.ID] = item
	}

	// 处理向量搜索结果
	for _, item := range vectorSearchRes.List {
		item.Score *= vectorWeight
		if existing, exists := dataMap[item.ID]; exists {
			// 合并分数，保留其他字段
			item.Score += existing.Score
		}
		dataMap[item.ID] = item
	}

	// 转换为切片并排序
	dataList := make([]ESSearchHitsItem, 0, len(dataMap))
	for _, item := range dataMap {
		dataList = append(dataList, item)
	}

	sort.Slice(dataList, func(i, j int) bool {
		return dataList[i].Score > dataList[j].Score
	})

	return &SearchResponse{
		Total:    len(dataList),
		MaxScore: dataList[0].Score,
		List:     dataList,
	}, nil
}

// hybridSearchByRRF 使用Reciprocal Rank Fusion的混合搜索
func hybridSearchByRRF(ctx *gin.Context, searchValue string) (*SearchResponse, error) {

	cfg := DefaultSearchConfig

	embedding, embeddingErr := singleEmbedding(ctx, searchValue)
	if embeddingErr != nil {
		return nil, fmt.Errorf("failed to generate embedding: %w", embeddingErr)
	}

	textQuery := dbes.BuildMap("standard",
		dbes.BuildMap("query", dbes.BuildMap("match", dbes.BuildMap("content", searchValue))))

	knnQuery := dbes.BuildMap("knn",
		dbes.BuildMap("field", "embedding",
			"query_vector", embedding,
			"k", cfg.K,
			"num_candidates", cfg.NumCandidates))
	queryBuilder := dbes.NewBuilder().Set("retriever",
		dbes.BuildMap("rrf",
			dbes.BuildMap("retrievers", []dbes.Map{textQuery, knnQuery}, "rank_window_size", 100, "rank_constant", 20))).
		SetSource([]string{"doc_id", "content", "category"})

	return executeSearch(ctx, queryBuilder)
}
