---
name: go-best-practices
description: Go 语言最佳实践和惯用模式，包括错误处理、并发、接口和项目结构指南。
---

# Go 最佳实践

## 错误处理
- 始终检查并返回错误；永远不要静默忽略。
- 使用 `fmt.Errorf("context: %w", err)` 为错误添加上下文。
- 使用 `var ErrXxx = errors.New(...)` 为导出的错误定义哨兵错误。
- 使用 `errors.Is` 和 `errors.As` 检查包装后的错误。

## 命名规范
- 局部变量使用简短、精炼的名称（如 `i`、`v`、`err`）。
- 导出的名称应当有意义且自文档化。
- 单方法接口名称应以 `-er` 结尾（如 `Reader`、`Writer`）。
- 避免冗余的包名（如使用 `log.Info` 而非 `log.LogInfo`）。

## 并发
- goroutine 之间通信优先使用 channel。
- 使用 `sync.WaitGroup` 等待一组 goroutine 完成。
- 使用 `sync.Mutex` 或 `sync.RWMutex` 保护共享状态。
- 始终将 context 作为第一个参数传递；通过取消 context 来停止 goroutine。

## 项目结构
- 保持 `main.go` 简洁；将业务逻辑放在包中。
- 将相关功能组织到子包中。
- 遵循标准布局：`cmd/`、`internal/`、`pkg/`。

## 测试
- 使用表驱动测试以获得全面的覆盖。
- 测试命名格式：`TestXxx_条件场景_预期结果`。
- 在测试辅助函数中使用 `t.Helper()`。
- 使用接口模拟外部依赖。