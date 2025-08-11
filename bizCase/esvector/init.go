package main

import (
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/morehao/golib/dbstore/dbes"
)

var ESClient *elasticsearch.Client

func init() {
	cfg := &dbes.ESConfig{
		Service: "es",
		Addr:    "http://localhost:9200",
	}
	simpleClient, _, err := dbes.InitES(cfg)
	if err != nil {
		panic(err)
	}
	ESClient = simpleClient
}

const (
	ESIndexName = "vector_research"

	EmbeddingHost = "xxx"

	SearchTypeText              = "text"
	SearchTypeVector            = "vector"
	SearchTypeTextVector        = "text_vector"
	SearchTypeHybridScriptScore = "hybrid_script_score"
	SearchTypeHybridMemoryScore = "hybrid_memory_score" // 内存中重新计算分数
	SearchTypeHybridRRF         = "hybrid_rrf"
)

type contentItem struct {
	Content  string `json:"content"`
	Category string `json:"category"`
}

var contentTestList = []contentItem{
	{
		Content:  "Elasticsearch 是一个基于 Lucene 的强大搜索引擎。",
		Category: "搜索引擎",
	},
	{
		Content:  "稠密向量可以将文本转化为数字表示，实现语义搜索。",
		Category: "向量检索",
	},
	{
		Content:  "敏捷的棕色狐狸跳过了懒狗。",
		Category: "通用测试",
	},
	{
		Content:  "机器学习可以帮助计算机从数据中学习规律。",
		Category: "人工智能",
	},
	{
		Content:  "旧金山的天气通常早上多雾。",
		Category: "地理天气",
	},
	{
		Content:  "混合检索结合了词法和语义两种搜索方式。",
		Category: "向量检索",
	},
	{
		Content:  "OpenAI 基于 GPT 架构开发了 ChatGPT。",
		Category: "人工智能",
	},
	{
		Content:  "Python 是一门广受欢迎的数据科学编程语言。",
		Category: "编程语言",
	},
	{
		Content:  "太阳能是最可持续的能源之一。",
		Category: "可再生能源",
	},
	{
		Content:  "篮球是一项在矩形球场上进行的团队运动。",
		Category: "体育运动",
	},
}
