package main

import (
	"fmt"

	"github.com/morehao/golib/database/dbredis"
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
