package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func ReadFromStream(ctx *gin.Context, streamKey, lastID string, count int64, block time.Duration) ([]redis.XMessage, error) {
	streams, err := rdb.XRead(ctx, &redis.XReadArgs{
		Streams: []string{streamKey, lastID},
		Count:   count,
		Block:   block,
	}).Result()
	if err != nil {
		// 例如 i/o timeout、context canceled 都会走这里
		return nil, err
	}
	if len(streams) == 0 {
		// 没有任何 Stream 返回
		return nil, nil
	}
	// 我们只关心第一个 Stream 的 Message（因为我们只传了一个 streamKey）
	return streams[0].Messages, nil
}

func SaveToStream(ctx *gin.Context, stream, data string) string {
	id, _ := rdb.XAdd(ctx, &redis.XAddArgs{
		Stream: stream,
		Values: map[string]interface{}{"data": data},
	}).Result()
	return id
}
