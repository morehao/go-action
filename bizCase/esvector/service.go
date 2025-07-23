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
			"num_candidates", cfg.NumCandidates,
			"boost", cfg.VectorWeight)).
		SetQuery(dbes.BuildMap("match", dbes.BuildMap(
			"content", dbes.BuildMap(
				"query", searchValue,
				"boost", cfg.TextWeight))))

	return executeSearch(ctx, queryBuilder)
}

// rrfSearch 使用Reciprocal Rank Fusion的混合搜索
func rrfSearch(ctx *gin.Context, searchValue string) (any, error) {

	cfg := DefaultSearchConfig

	// 并行执行文本搜索和向量搜索
	textResultChan := make(chan any, 1)
	vectorResultChan := make(chan any, 1)
	textErrChan := make(chan error, 1)
	vectorErrChan := make(chan error, 1)

	// 执行文本搜索
	go func() {
		result, err := textSearch(ctx, searchValue)
		textResultChan <- result
		textErrChan <- err
	}()

	// 执行向量搜索
	go func() {
		result, err := vectorSearch(ctx, searchValue)
		vectorResultChan <- result
		vectorErrChan <- err
	}()

	// 等待两个搜索完成
	textResult := <-textResultChan
	vectorResult := <-vectorResultChan
	textErr := <-textErrChan
	vectorErr := <-vectorErrChan

	if textErr != nil {
		return nil, fmt.Errorf("text search failed: %w", textErr)
	}

	if vectorErr != nil {
		return nil, fmt.Errorf("vector search failed: %w", vectorErr)
	}

	// 合并结果使用RRF
	return combineResultsWithRRF(textResult, vectorResult, cfg.RRFConstant, cfg.K)
}

// combineResultsWithRRF 使用RRF算法合并搜索结果
func combineResultsWithRRF(textResult, vectorResult any, rrfConstant, topK int) (map[string]any, error) {
	textHits, err := extractHits(textResult)
	if err != nil {
		return nil, fmt.Errorf("failed to extract text search hits: %w", err)
	}

	vectorHits, err := extractHits(vectorResult)
	if err != nil {
		return nil, fmt.Errorf("failed to extract vector search hits: %w", err)
	}

	// 文档ID到排名的映射
	docRanks := make(map[string]map[string]int)
	allDocs := make(map[string]*SearchResult)

	// 处理文本搜索结果
	for rank, hit := range textHits {
		docID := getDocID(hit)
		if docID != "" {
			if docRanks[docID] == nil {
				docRanks[docID] = make(map[string]int)
			}
			docRanks[docID]["text"] = rank + 1

			if _, exists := allDocs[docID]; !exists {
				allDocs[docID] = hitToSearchResult(hit)
			}
		}
	}

	// 处理向量搜索结果
	for rank, hit := range vectorHits {
		docID := getDocID(hit)
		if docID != "" {
			if docRanks[docID] == nil {
				docRanks[docID] = make(map[string]int)
			}
			docRanks[docID]["vector"] = rank + 1

			if _, exists := allDocs[docID]; !exists {
				allDocs[docID] = hitToSearchResult(hit)
			}
		}
	}

	// 计算RRF分数
	type docScore struct {
		docID string
		score float64
		doc   *SearchResult
	}

	var docScores []docScore
	for docID, ranks := range docRanks {
		rrfScore := 0.0
		if textRank, hasText := ranks["text"]; hasText {
			rrfScore += 1.0 / float64(textRank+rrfConstant)
		}
		if vectorRank, hasVector := ranks["vector"]; hasVector {
			rrfScore += 1.0 / float64(vectorRank+rrfConstant)
		}

		doc := allDocs[docID]
		doc.Score = rrfScore

		docScores = append(docScores, docScore{
			docID: docID,
			score: rrfScore,
			doc:   doc,
		})
	}

	// 按RRF分数排序
	sort.Slice(docScores, func(i, j int) bool {
		return docScores[i].score > docScores[j].score
	})

	// 取前topK个结果
	if len(docScores) > topK {
		docScores = docScores[:topK]
	}

	// 构造返回结果
	hits := make([]map[string]any, len(docScores))
	for i, ds := range docScores {
		hits[i] = map[string]any{
			"_id":    ds.docID,
			"_score": ds.score,
			"_source": map[string]any{
				"doc_id":   ds.doc.DocID,
				"content":  ds.doc.Content,
				"category": ds.doc.Category,
			},
		}
	}

	return map[string]any{
		"hits": map[string]any{
			"total": map[string]any{
				"value": len(hits),
			},
			"hits": hits,
		},
	}, nil
}

// 辅助函数
func extractHits(result any) ([]map[string]any, error) {
	resultMap, ok := result.(map[string]any)
	if !ok {
		return nil, fmt.Errorf("invalid result format")
	}

	hits, ok := resultMap["hits"].(map[string]any)
	if !ok {
		return nil, fmt.Errorf("no hits found in result")
	}

	hitsList, ok := hits["hits"].([]any)
	if !ok {
		return nil, fmt.Errorf("invalid hits format")
	}

	var result_hits []map[string]any
	for _, hit := range hitsList {
		if hitMap, ok := hit.(map[string]any); ok {
			result_hits = append(result_hits, hitMap)
		}
	}

	return result_hits, nil
}

func getDocID(hit map[string]any) string {
	if source, ok := hit["_source"].(map[string]any); ok {
		if docID, ok := source["doc_id"].(string); ok {
			return docID
		}
	}
	if id, ok := hit["_id"].(string); ok {
		return id
	}
	return ""
}

func hitToSearchResult(hit map[string]any) *SearchResult {
	result := &SearchResult{}

	if source, ok := hit["_source"].(map[string]any); ok {
		if docID, ok := source["doc_id"].(string); ok {
			result.DocID = docID
		}
		if content, ok := source["content"].(string); ok {
			result.Content = content
		}
		if category, ok := source["category"].(string); ok {
			result.Category = category
		}
		result.Source = source
	}

	if score, ok := hit["_score"].(float64); ok {
		result.Score = score
	}

	return result
}
