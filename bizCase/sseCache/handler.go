package main

import (
	"io"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/morehao/golib/glog"
	"github.com/ygpkg/yg-go/apis/ssecache"
)

func Chat(ctx *gin.Context) {
	questionID := "2025"
	ch := make(chan string)
	sseCache := ssecache.NewSSEClient(ssecache.WithChannel(ch), ssecache.WithRedisClient(rdb))
	var writeCount int
	var mu sync.Mutex

	go func() {
		ticker := time.NewTicker(200 * time.Millisecond)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				mu.Lock()
				if writeCount > 100 {
					mu.Unlock()
					return
				}
				msg := time.Now().Format(time.RFC3339)
				if err := sseCache.WriteMessage(ctx, questionID, msg); err != nil {
					glog.Errorf(ctx, "[Chat] WriteMessage failed: %v", err)
					mu.Unlock()
					return
				}
				writeCount++
				mu.Unlock()
			}
		}
	}()

	ctx.Stream(func(w io.Writer) bool {
		stoped, err := sseCache.GetStopSignal(ctx, questionID)
		if err != nil || stoped {
			return false
		}

		select {
		case msg := <-ch:
			ctx.SSEvent("message", msg)
		case <-time.After(100 * time.Millisecond):
		}

		return true
	})

	// TODO：结束后的操作，关闭 channel、删除 redis？
}

func StopChat(ctx *gin.Context) {
	questionID := "2025"
	sseCache := ssecache.NewSSEClient(ssecache.WithRedisClient(rdb))
	if err := sseCache.Stop(ctx, questionID); err != nil {
		glog.Errorf(ctx, "[StopChat] sseCache.Stop failed, err: %v", err)
	}
	glog.Infof(ctx, "[StopChat] completed")
}

func GetMessage(ctx *gin.Context) {
	questionID := "2025"
	sseCache := ssecache.NewSSEClient(ssecache.WithRedisClient(rdb))
	messages, err := sseCache.ReadMessages(ctx, questionID, "")
	if err != nil {
		glog.Errorf(ctx, "[GetMessage] sseCache.ReadMessages failed, err: %v", err)
	} else {
		ctx.JSON(200, messages)
	}
}
