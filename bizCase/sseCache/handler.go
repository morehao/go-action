package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/morehao/golib/glog"
	"github.com/morehao/golib/gutils"
	"github.com/ygpkg/yg-go/apis/sseclient"
)

func Chat(ctx *gin.Context) {
	questionID := "202506261643"
	sseClient := sseclient.New(sseclient.WithRedisClient(rdb))
	sseClient.SetHeaders(ctx.Writer)
	var writeCount int
	var mu sync.Mutex
	for i := 0; i < 100; i++ {
		msg := time.Now().Format(time.RFC3339)
		msg = fmt.Sprintf("id: %d, message: %s\n", writeCount, msg)
		if stoped, err := sseClient.WriteMessage(ctx, ctx.Writer, questionID, msg); err != nil {
			glog.Errorf(ctx, "[Chat] WriteMessage failed: %v", err)
			mu.Unlock()
			return
		} else if stoped {
			mu.Unlock()
			return
		}
		writeCount++
		time.Sleep(500 * time.Millisecond)
	}
}

func StopChat(ctx *gin.Context) {
	questionID := "202506261643"
	sseClient := sseclient.New(sseclient.WithRedisClient(rdb))
	if err := sseClient.Stop(ctx, questionID); err != nil {
		glog.Errorf(ctx, "[StopChat] sseClient.Stop failed, err: %v", err)
	}
	glog.Infof(ctx, "[StopChat] completed")
}

func GetMessage(ctx *gin.Context) {
	questionID := "202506261643"
	sseClient := sseclient.New(sseclient.WithRedisClient(rdb))
	latestID, historyMessages, getHistoryMessageErr := sseClient.ReadMessages(ctx, questionID)
	if getHistoryMessageErr != nil {
		glog.Errorf(ctx, "[GetMessage] sseClient.ReadMessages failed, err: %v", getHistoryMessageErr)
	}
	glog.Infof(ctx, "[GetMessage] latestID: %s, historyMessages: %v", latestID, gutils.ToJsonString(historyMessages))
	ctx.SSEvent("history", gutils.ToJsonString(historyMessages))
	ctx.Writer.Flush()
	done, affectedRaw, readErr := sseClient.BlockRead(ctx, ctx.Writer, questionID, latestID)
	if readErr != nil {
		glog.Errorf(ctx, "[GetMessage] sseClient.BlockRead failed, err: %v", readErr)
	}
	glog.Infof(ctx, "[GetMessage] affectedRaw: %d, done: %v", affectedRaw, done)
	glog.Infof(ctx, "[GetMessage] completed")

}
