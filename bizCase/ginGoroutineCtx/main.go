package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/test", func(ctx *gin.Context) {
		// 往上下文中写入数据
		unixMilli := time.Now().UnixMilli()
		ctx.Set("timestamp", unixMilli)

		// 在主线程中启动一个 goroutine
		go func() {
			// 模拟耗时任务
			time.Sleep(10 * time.Second)

			// 从上下文中读取数据并比较
			value, exists := ctx.Get("timestamp")
			if exists {
				// 比较时间戳
				if value.(int64) == unixMilli {
					println("时间戳相同")
				} else {
					println("时间戳不同")
				}
			} else {
				println("数据不存在")
			}
		}()

		ctx.JSON(200, gin.H{
			"message": "test success",
		})
	})

	r.GET("testfix", func(ctx *gin.Context) {
		// 往上下文中写入数据
		unixMilli := time.Now().UnixMilli()
		ctx.Set("timestamp", unixMilli)

		// 使用 ctx 的副本
		goroutineCtx := ctx.Copy()
		// 在主线程中启动一个 goroutine
		go func() {
			// 模拟耗时任务
			time.Sleep(10 * time.Second)

			// 从上下文中读取数据并比较
			value, exists := goroutineCtx.Get("timestamp")
			if exists {
				// 比较时间戳
				if value.(int64) == unixMilli {
					println("时间戳相同")
				} else {
					println("时间戳不同")
				}
			} else {
				println("数据不存在")
			}
		}()

		ctx.JSON(200, gin.H{
			"message": "test success",
		})
	})

	r.GET("/healthcheck", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "ok",
		})
	})

	r.Run(":8080")
}
