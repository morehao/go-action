package main

import (
	"encoding/json"
	"fmt"

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
func executeSearch(ctx *gin.Context, queryBuilder *dbes.Builder) (*map[string]any, error) {
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

	var res map[string]any
	if err := json.NewDecoder(searchRes.Body).Decode(&res); err != nil {
		return nil, fmt.Errorf("failed to decode search response: %w", err)
	}

	return &res, nil
}

// textSearch 文本搜索
func textSearch(ctx *gin.Context, searchValue string) (any, error) {

	cfg := DefaultSearchConfig

	queryBuilder := dbes.NewBuilder().
		SetSource([]string{"doc_id", "content", "category"}).
		SetSize(cfg.K).
		SetQuery(dbes.BuildMap("match", dbes.BuildMap("content", searchValue)))

	return executeSearch(ctx, queryBuilder)
}

// vectorSearch 向量搜索
func vectorSearch(ctx *gin.Context, searchValue string) (any, error) {

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

// hybridSearch 真正的混合搜索（使用Elasticsearch的原生混合搜索）
func hybridSearch(ctx *gin.Context, searchValue string) (any, error) {

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

// rrfSearch 使用Reciprocal Rank Fusion的混合搜索
func rrfSearch(ctx *gin.Context, searchValue string) (any, error) {

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
