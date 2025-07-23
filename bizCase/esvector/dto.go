package main

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

type SearchRequest struct {
	SearchValue string `json:"searchValue" validate:"required"` // 检索词
	SearchType  string `json:"searchType" validate:"required"`  // 检索类型，text：文本检索，vector：向量检索，hybrid：混合检索
}
