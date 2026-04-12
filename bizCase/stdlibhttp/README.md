# stdlibhttp - Go 标准库 net/http Web 服务示例

基于 Go 标准库 `net/http` 构建的 Web 服务示例，不依赖任何第三方框架。

## 功能特性

- **原生路由**: 使用 `http.ServeMux` 实现 RESTful 路由
- **中间件支持**: 日志、错误恢复、CORS
- **JSON 处理**: 统一的 JSON 响应格式
- **优雅关闭**: 信号中断处理

## 运行

```bash
go run .
```

## API 接口

### 健康检查
```bash
curl http://localhost:8080/health
```

### Hello 接口
```bash
curl http://localhost:8080/hello?name=Go
```

### 获取用户列表
```bash
curl http://localhost:8080/users
```

### 创建用户
```bash
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"name":"Charlie","email":"charlie@example.com"}'
```