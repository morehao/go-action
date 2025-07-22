package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/morehao/golib/protocol"
	"github.com/morehao/golib/protocol/gresty"
)

type EmbeddingRequest struct {
	Model string `json:"model"`
	Input string `json:"input"`
}

type EmbeddingResponse struct {
	Data []EmbeddingDataItem `json:"data"`
}

type EmbeddingDataItem struct {
	Embedding []float32 `json:"embedding"`
}

func singleEmbedding(ctx *gin.Context, content string) (embedding []float32, err error) {
	httpClientCfg := &protocol.HttpClientConfig{
		Module: "esvector",
		Host:   "https://xxxxx",
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
