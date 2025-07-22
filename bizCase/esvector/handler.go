package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/morehao/golib/gcontext/gincontext"
	"github.com/morehao/golib/glog"
	"github.com/morehao/golib/storages/dbes"
)

func PutVector(ctx *gin.Context) {
	text := "这是一个用于测试的文本"
	embedding := make([]float32, 384)
	for i := range embedding {
		embedding[i] = float32(i) * 0.001
	}

	doc := map[string]any{
		"content":   text,
		"embedding": embedding,
	}
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(doc); err != nil {
		glog.Errorf(ctx, "[PutVector] json.NewEncoder.Encode error: %v", err)
		gincontext.Fail(ctx, err)
		return
	}

	res, err := ESClient.Index("vector_test", &buf)
	if err != nil {
		glog.Errorf(ctx, "[PutVector] ESClient.Index error: %v", err)
		gincontext.Fail(ctx, err)
		return
	}
	defer res.Body.Close()
	gincontext.Success(ctx, res.String())
}

func GetVector(ctx *gin.Context) {
	embedding := make([]float32, 384)
	for i := range embedding {
		embedding[i] = float32(i) * 0.001
	}
	queryBuilder := dbes.NewBuilder().SetSource([]string{"content"}).
		Set("knn", dbes.BuildMap("field", "embedding", "query_vector", embedding, "k", 5, "num_candidates", 100))

	buf, err := queryBuilder.BuildReader()
	if err != nil {
		glog.Errorf(ctx, "[GetVector] queryBuilder.BuildReader error: %v", err)
		gincontext.Fail(ctx, err)
		return
	}

	res, err := ESClient.Search(
		ESClient.Search.WithContext(context.Background()),
		ESClient.Search.WithIndex("vector_test"),
		ESClient.Search.WithBody(buf),
	)
	if err != nil {
		glog.Errorf(ctx, "[GetVector] ESClient.Search error: %v", err)
		gincontext.Fail(ctx, err)
		return
	}
	if res.StatusCode != 200 {
		glog.Errorf(ctx, "[GetVector] ESClient.Search error: %v", res.Status)
		gincontext.Fail(ctx, fmt.Errorf("%s", res.Status()))
		return
	}
	defer res.Body.Close()

	var r map[string]any
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		glog.Errorf(ctx, "[GetVector] json.NewDecoder.Decode error: %v", err)
		gincontext.Fail(ctx, err)
		return
	}

	for _, hit := range r["hits"].(map[string]any)["hits"].([]any) {
		source := hit.(map[string]any)["_source"]
		fmt.Printf("匹配内容: %+v\n", source)
	}
	gincontext.Success(ctx, r)
}
