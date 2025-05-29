package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/morehao/golib/storages/dbredis"
	"github.com/redis/go-redis/v9"
)

var (
	rdb *redis.Client
)

// 初始化 Redis 客户端
func init() {
	cfg := dbredis.RedisConfig{
		Service: "redis",
		Addr:    "127.0.0.1:6379",
		DB:      0,
	}
	redisClient, initErr := dbredis.InitRedis(&cfg)
	if initErr != nil {
		panic(fmt.Sprintf("初始化 Redis 失败: %v", initErr))
	}
	rdb = redisClient
}

func ReadFromStream(ctx *gin.Context, stream, lastID string, count int64, block time.Duration) ([]redis.XMessage, error) {
	res, err := rdb.XRead(ctx, &redis.XReadArgs{
		Streams: []string{stream, lastID},
		Count:   count,
		Block:   block,
	}).Result()
	if err != nil || len(res) == 0 {
		return nil, err
	}
	return res[0].Messages, nil
}

func SaveToStream(ctx *gin.Context, stream, data string) string {
	id, _ := rdb.XAdd(ctx, &redis.XAddArgs{
		Stream: stream,
		Values: map[string]interface{}{"data": data},
	}).Result()
	return id
}
