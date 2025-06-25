package main

import (
	"github.com/go-redis/redis/v8"
)

var (
	rdb *redis.Client
)

// 初始化 Redis 客户端
func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		DB:   0,
	})
}
