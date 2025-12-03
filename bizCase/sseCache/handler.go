/*
 * @Author: morehao morehao@qq.com
 * @Date: 2025-07-22 14:36:10
 * @LastEditors: morehao morehao@qq.com
 * @LastEditTime: 2025-12-02 20:12:10
 * @FilePath: /go-action/bizCase/sseCache/handler.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/morehao/golib/glog"
	"github.com/morehao/golib/gutil"
	"github.com/ygpkg/yg-go/apis/sseclient"
)

func Chat(ctx *gin.Context) {
	questionID := "202506261643"
	sseClient := sseclient.New(sseclient.WithRedisClient(rdb), sseclient.WithExpiration(time.Minute*5))
	sseClient.SetHeaders(ctx.Writer)
	var writeCount int
	for i := 0; i < 100; i++ {
		msg := time.Now().Format(time.RFC3339)
		msg = fmt.Sprintf("id: %d, message: %s\n", writeCount, msg)
		if stoped, err := sseClient.WriteMessage(ctx, ctx.Writer, questionID, msg); err != nil {
			glog.Errorf(ctx, "[Chat] WriteMessage failed: %v", err)
			return
		} else if stoped {
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
	glog.Infof(ctx, "[GetMessage] latestID: %s, historyMessages: %v", latestID, gutil.ToJsonString(historyMessages))

	// 发送历史消息
	if err := sseClient.SendEvent(ctx.Writer, fmt.Sprintf("history: %s\n", gutil.ToJsonString(historyMessages))); err != nil {
		glog.Errorf(ctx, "[GetMessage] sseClient.SendEvent failed, err: %v", err)
	}

	// 阻塞读取剩余消息
	done, affectedRaw, readErr := sseClient.BlockRead(ctx, ctx.Writer, questionID, latestID)
	if readErr != nil {
		glog.Errorf(ctx, "[GetMessage] sseClient.BlockRead failed, err: %v", readErr)
	}
	glog.Infof(ctx, "[GetMessage] affectedRaw: %d, done: %v", affectedRaw, done)
	glog.Infof(ctx, "[GetMessage] completed")

}
