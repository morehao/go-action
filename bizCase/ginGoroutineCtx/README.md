# 问题描述
  当在中间件或 handler 中启动新的 Goroutine 时，不能使用原始的上下文，必须使用只读副本
# 场景复现
1. 启动示例项目: `go run main.go`
2. 构造请求目标列表`targets.txt`，内容如下:
    ```
    GET http://localhost:8080/testfix
    GET http://localhost:8080/healthcheck
    ```
3. 使用 vegeta 进行压测, 每秒发送 10 个请求，持续 1s，命令如下
    ```
    vegeta attack -rate=10 -duration=1s -targets=targets.txt | vegeta report
    ```
4. 查看输出，如下：
    ``` bash
    [GIN] 2025/06/18 - 16:21:06 | 200 |     124.292µs |       127.0.0.1 | GET      "/test"
    [GIN] 2025/06/18 - 16:21:06 | 200 |       18.75µs |       127.0.0.1 | GET      "/healthcheck"
    [GIN] 2025/06/18 - 16:21:06 | 200 |          58µs |       127.0.0.1 | GET      "/test"
    [GIN] 2025/06/18 - 16:21:06 | 200 |      30.084µs |       127.0.0.1 | GET      "/healthcheck"
    [GIN] 2025/06/18 - 16:21:06 | 200 |      70.334µs |       127.0.0.1 | GET      "/test"
    [GIN] 2025/06/18 - 16:21:06 | 200 |      18.584µs |       127.0.0.1 | GET      "/healthcheck"
    [GIN] 2025/06/18 - 16:21:06 | 200 |       149.5µs |       127.0.0.1 | GET      "/test"
    [GIN] 2025/06/18 - 16:21:06 | 200 |      29.667µs |       127.0.0.1 | GET      "/healthcheck"
    [GIN] 2025/06/18 - 16:21:07 | 200 |      43.625µs |       127.0.0.1 | GET      "/test"
    [GIN] 2025/06/18 - 16:21:07 | 200 |      16.833µs |       127.0.0.1 | GET      "/healthcheck"
    时间戳不同
    数据不存在
    时间戳不同
    时间戳相同
    时间戳相同
    ```
预期控制台打印信息应始终为：时间戳相同，但实际情况却还出现：
- 时间戳不同
- 数据不存在

# 原因分析
输出不符合预期，是因为`key`为`timestamp`的数据被删除。

`gin.Context`相关源码如下：
``` go
// ServeHTTP conforms to the http.Handler interface.
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := engine.pool.Get().(*Context)
	c.writermem.reset(w)
	c.Request = req
	c.reset()

	engine.handleHTTPRequest(c)

	engine.pool.Put(c)
}


func (c *Context) reset() {
	c.Writer = &c.writermem
	c.Params = c.Params[:0]
	c.handlers = nil
	c.index = -1

	c.fullPath = ""
	c.Keys = nil
	c.Errors = c.Errors[:0]
	c.Accepted = nil
	c.queryCache = nil
	c.formCache = nil
	c.sameSite = 0
	*c.params = (*c.params)[:0]
	*c.skippedNodes = (*c.skippedNodes)[:0]
}
```
分析如下：
- `ServeHTTP` 是 HTTP 请求的入口方法
- `ServeHTTP` 中使用对象池复用 `Context` 对象，并非每次请求都新建 `Context`
- 获取到 `Context` 对象后，会通过 `reset` 方法清空状态和数据
- 在 `/test` 接口中，返回响应后，`Context` 对象会被放回对象池。而 `Goroutine` 延迟10 秒后才从 `Context` 对象中读取 `timestamp` 的数据。在这10 秒里，`Context` 对象可能已经被复用到其他请求 
  - 复用到 `/test` 接口请求里，导致timestamp 的值被覆盖。
  - 复用到 `/healthcheck` 接口请求里，导致 `timestamp` 被删除（因为 `reset` 清空了数据）。

# 解决方案
`Gin` 框架提供了 `context.Copy()` 方法，用于创建上下文的只读副本。副本是协程安全的，因为它复制了上下文中的大部分数据，同时与原始上下文隔离。修改后的`/test`接口如下：
``` go
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
```