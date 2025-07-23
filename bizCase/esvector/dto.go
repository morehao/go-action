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

// SearchConfig 搜索配置
type SearchConfig struct {
	K             int     `json:"k"`
	NumCandidates int     `json:"num_candidates"`
	RRFConstant   int     `json:"rrf_constant"`
	TextWeight    float64 `json:"text_weight"`
	VectorWeight  float64 `json:"vector_weight"`
}

// DefaultSearchConfig 默认搜索配置
var DefaultSearchConfig = SearchConfig{
	K:             10,
	NumCandidates: 100,
	RRFConstant:   60,
	TextWeight:    0.5,
	VectorWeight:  0.5,
}

// ESSearchResponse 标准化的搜索结果
type ESSearchResponse struct {
	Hits    ESSearchHits  `json:"hits"`
	TimeOut bool          `json:"time_out"`
	Took    int           `json:"took"`
	Shard   ESSearchShard `json:"_shards"`
}

type ESSearchShard struct {
	Total      int `json:"total"`
	Successful int `json:"successful"`
	Failed     int `json:"failed"`
	Skipped    int `json:"skipped"`
}

type ESSearchHits struct {
	Hits     []ESSearchHitsItem `json:"hits"`
	MaxScore float64            `json:"max_score"`
	Total    ESSearchHitsTotal  `json:"total"`
}

type ESSearchHitsTotal struct {
	Relation string `json:"relation"`
	Value    int    `json:"value"`
}

type ESSearchHitsItem struct {
	ID     string                 `json:"_id"`
	Index  string                 `json:"_index"`
	Score  float64                `json:"_score"`
	Source ESSearchHitsItemSource `json:"_source"`
}

type ESSearchHitsItemSource struct {
	DocID    uint   `json:"doc_id"`
	Content  string `json:"content"`
	Category string `json:"category"`
}

type SearchResponse struct {
	Total    int                `json:"total"`
	MaxScore float64            `json:"max_score"`
	List     []ESSearchHitsItem `json:"list"`
}
