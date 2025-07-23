package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/morehao/golib/gcontext/gincontext"
	"github.com/morehao/golib/glog"
)

func InsertData(ctx *gin.Context) {

	for _, v := range contentTestList {

		embedding, embeddingErr := singleEmbedding(ctx, v.Content)
		if embeddingErr != nil {
			glog.Errorf(ctx, "[InsertData] singleEmbedding error: %v", embeddingErr)
			gincontext.Fail(ctx, embeddingErr)
			return
		}
		doc := map[string]any{
			"doc_id":    time.Now().UnixNano(),
			"content":   v.Content,
			"embedding": embedding,
			"category":  v.Category,
		}
		var buf bytes.Buffer
		if err := json.NewEncoder(&buf).Encode(doc); err != nil {
			glog.Errorf(ctx, "[InsertData] json.NewEncoder.Encode error: %v", err)
			gincontext.Fail(ctx, err)
			return
		}
		res, err := ESClient.Index(ESIndexName, &buf)
		if err != nil {
			glog.Errorf(ctx, "[InsertData] ESClient.Index error: %v", err)
			gincontext.Fail(ctx, err)
			return
		}
		defer res.Body.Close()
		fmt.Printf("res: %s\n", res.String())

	}
	gincontext.Success(ctx, "success")
}

func SearchData(ctx *gin.Context) {
	// 从 request 的 body 中获取搜索词
	var req SearchRequest
	if err := ctx.BindJSON(&req); err != nil {
		glog.Errorf(ctx, "[SearchData] ctx.BindJSON error: %v", err)
		gincontext.Fail(ctx, err)
		return
	}

	var searchRes any
	var searchErr error

	switch req.SearchType {
	case SearchTypeText:
		searchRes, searchErr = textSearch(ctx, req.SearchValue)
	case SearchTypeVector:
		searchRes, searchErr = vectorSearch(ctx, req.SearchValue)
	case SearchTypeHybrid:
		searchRes, searchErr = hybridSearch(ctx, req.SearchValue)
	default:
		searchRes, searchErr = textSearch(ctx, req.SearchValue)
	}

	if searchErr != nil {
		glog.Errorf(ctx, "[SearchData] search error: %v", searchErr)
		gincontext.Fail(ctx, searchErr)
		return
	}
	gincontext.Success(ctx, searchRes)
}
