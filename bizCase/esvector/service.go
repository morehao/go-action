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
		Host:   "xxxx",
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

func textSearch(ctx *gin.Context, searchValue string) (any, error) {
	queryBuilder := dbes.NewBuilder().SetSource([]string{"doc_id", "content", "category"}).
		SetQuery(dbes.BuildMap("match", dbes.BuildMap("content", searchValue)))
	searchBody, buildErr := queryBuilder.BuildReader()
	if buildErr != nil {
		return nil, buildErr
	}
	searchRes, searchErr := ESClient.Search(
		ESClient.Search.WithContext(ctx),
		ESClient.Search.WithIndex(ESIndexName),
		ESClient.Search.WithBody(searchBody))
	if searchErr != nil {
		return nil, searchErr
	}

	defer searchRes.Body.Close()

	if searchRes.StatusCode != 200 {
		return nil, fmt.Errorf("%s", searchRes.Status())
	}
	var res map[string]any
	if err := json.NewDecoder(searchRes.Body).Decode(&res); err != nil {
		return nil, err
	}
	return res, nil
}

func vectorSearch(ctx *gin.Context, searchValue string) (any, error) {
	embedding, embeddingErr := singleEmbedding(ctx, searchValue)
	if embeddingErr != nil {
		return nil, embeddingErr
	}
	queryBuilder := dbes.NewBuilder().SetSource([]string{"doc_id", "content", "category"}).
		Set("knn", dbes.BuildMap("field", "embedding", "query_vector", embedding, "k", 10, "num_candidates", 100))
	searchBody, buildErr := queryBuilder.BuildReader()
	if buildErr != nil {
		return nil, buildErr
	}
	searchRes, searchErr := ESClient.Search(
		ESClient.Search.WithContext(ctx),
		ESClient.Search.WithIndex(ESIndexName),
		ESClient.Search.WithBody(searchBody))
	if searchErr != nil {
		return nil, searchErr
	}

	defer searchRes.Body.Close()

	if searchRes.StatusCode != 200 {
		return nil, fmt.Errorf("%s", searchRes.Status())
	}

	var res map[string]any
	if err := json.NewDecoder(searchRes.Body).Decode(&res); err != nil {
		return nil, err
	}
	return res, nil
}

func hybridSearch(ctx *gin.Context, searchValue string) (any, error) {
	embedding, embeddingErr := singleEmbedding(ctx, searchValue)
	if embeddingErr != nil {
		return nil, embeddingErr
	}
	queryBuilder := dbes.NewBuilder().SetSource([]string{"doc_id", "content", "category"}).
		Set("knn", dbes.BuildMap("field", "embedding", "query_vector", embedding, "k", 10, "num_candidates", 100)).
		SetQuery(dbes.BuildMap("match", dbes.BuildMap("content", searchValue)))
	searchBody, buildErr := queryBuilder.BuildReader()
	if buildErr != nil {
		return nil, buildErr
	}
	searchRes, searchErr := ESClient.Search(
		ESClient.Search.WithContext(ctx),
		ESClient.Search.WithIndex(ESIndexName),
		ESClient.Search.WithBody(searchBody))
	if searchErr != nil {
		return nil, searchErr
	}

	defer searchRes.Body.Close()

	if searchRes.StatusCode != 200 {
		return nil, fmt.Errorf("%s", searchRes.Status())
	}

	var res map[string]any
	if err := json.NewDecoder(searchRes.Body).Decode(&res); err != nil {
		return nil, err
	}
	return res, nil
}
