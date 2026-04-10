# Excel Agent HTTP Service 设计文档

## 1. 项目概述

- **项目名称**: excelagent
- **项目类型**: 基于 Gin + Eino 的 HTTP 服务
- **核心功能**: 提供 Excel/CSV 数据处理的 RESTful API 服务，采用 Plan-Execute-Replan 模式
- **目标用户**: 需要通过 HTTP 接口处理 Excel/CSV 数据的开发者

## 2. 架构设计

### 2.1 整体架构

```
┌─────────────────────────────────────────────────────────────────┐
│                         HTTP Client                             │
└─────────────────────────────────────────────────────────────────┘
                               │
                               ▼
┌─────────────────────────────────────────────────────────────────┐
│                      Gin HTTP Server                            │
│  ┌──────────────────────────────────────────────────────────┐  │
│  │                   POST /api/v1/tasks                      │  │
│  │            (同步执行，直接返回完整结果)                   │  │
│  └──────────────────────────────────────────────────────────┘  │
└─────────────────────────────────────────────────────────────────┘
                               │
                               ▼
┌─────────────────────────────────────────────────────────────────┐
│                    Eino AI Processing Pipeline                   │
│  ┌──────────────────────────────────────────────────────────┐  │
│  │  PlanExecuteAgent (Plan-Execute-Replan)                   │  │
│  │  ├── Planner: 解析需求，生成步骤计划                       │  │
│  │  ├── Executor: 执行每一步骤 (CodeAgent)                   │  │
│  │  └── Replanner: 决策继续/重规划/完成                       │  │
│  └──────────────────────────────────────────────────────────┘  │
└─────────────────────────────────────────────────────────────────┘
```

### 2.2 组件说明

| 组件 | 职责 |
|------|------|
| Gin Server | 提供 RESTful API 接口 |
| Eino Agent Pipeline | 核心 AI 处理逻辑，复用 excel-agent 的设计 |

## 3. API 设计

### 3.1 接口列表

| 方法 | 路径 | 描述 | 请求体 | 响应 |
|------|------|------|--------|------|
| POST | /api/v1/tasks | 创建并同步执行任务 | TaskRequest | TaskResultResponse |

### 3.2 数据模型

```go
// TaskRequest 创建任务请求
type TaskRequest struct {
    Query          string    `json:"query" binding:"required"` // 用户需求描述
    Files          []File    `json:"files"`                   // 上传的文件列表
    Options        *TaskOptions `json:"options,omitempty"`    // 可选配置
}

// File 文件信息
type File struct {
    Name string `json:"name"` // 文件名
    Type string `json:"type"` // 文件类型: xlsx, csv
}

// TaskOptions 任务选项
type TaskOptions struct {
    MaxIterations int `json:"max_iterations,omitempty"` // 最大迭代次数
    Timeout       int `json:"timeout,omitempty"`        // 超时时间(秒)
}

// TaskResultResponse 任务结果响应
type TaskResultResponse struct {
    TaskID      string    `json:"task_id"`
    Status      string    `json:"status"`         // completed / failed
    Result      string    `json:"result"`         // 处理结果描述
    OutputFiles []string  `json:"output_files"`  // 输出的文件列表
    Report      string    `json:"report"`         // 生成的报告
    CompletedAt time.Time `json:"completed_at"`
}
```

### 3.3 API 示例

**创建并执行任务**
```bash
# 请求
curl -X POST http://localhost:8080/api/v1/tasks \
  -H "Content-Type: application/json" \
  -d '{
    "query": "统计文件中推荐的小说名称及推荐次数，表头为小说名称和推荐次数",
    "files": [{"name": "data.xlsx", "type": "xlsx"}]
  }'

# 响应
{
  "task_id": "uuid-xxxx-xxxx",
  "status": "completed",
  "result": "处理完成，已生成统计结果",
  "output_files": ["result.csv"],
  "report": "## Task Report\n\n### Result\n统计完成...\n\n### Output Files\n- result.csv\n",
  "completed_at": "2025-01-01T10:01:00Z"
}
```

**任务执行失败**
```json
{
  "task_id": "uuid-xxxx-xxxx",
  "status": "failed",
  "result": "Error: file not found",
  "output_files": [],
  "report": "## Task Report\n\n### Result\nError: file not found",
  "completed_at": "2025-01-01T10:01:00Z"
}
```

## 4. 核心模块设计

### 4.1 项目结构

```
excelagent/
├── main.go                 # 程序入口
├── config/
│   └── config.go           # 配置管理
├── handler/
│   └── task_handler.go     # HTTP 处理器（同步执行）
├── service/
│   └── agent/
│       ├── agent.go        # Agent 封装
│       ├── planner.go      # Planner
│       ├── executor.go     # Executor
│       ├── replanner.go    # Replanner
│       ├── code_agent.go   # CodeAgent
│       └── report.go       # ReportAgent
├── tools/                  # 工具封装
│   └── ...
├── types/
│   └── types.go            # 类型定义
└── go.mod
```

### 4.2 同步任务处理

- 任务提交后同步执行 Eino Agent
- 处理完成后直接返回完整结果
- 无需任务状态查询和轮询

## 5. 复用设计

### 5.1 复用的组件

| 组件 | 复用方式 |
|------|---------|
| Planner | 直接复用 excel-agent 的实现 |
| Executor | 直接复用 excel-agent 的实现 |
| Replanner | 直接复用 excel-agent 的实现 |
| CodeAgent | 直接复用 excel-agent 的实现 |
| ReportAgent | 直接复用 excel-agent 的实现 |
| Tools | 直接复用 excel-agent 的工具 |

### 5.2 需要适配的部分

1. **文件上传处理**: 通过 HTTP 上传文件，保存到临时目录
2. **Context 传递**: 通过 gin.Context 传递任务信息到 Agent
3. **结果输出**: 适配 HTTP 响应格式

## 6. 配置说明

```yaml
# config.yaml
server:
  host: "0.0.0.0"
  port: 8080

eino:
  api_key: "${EINO_API_KEY}"
  base_url: "${EINO_BASE_URL}"
  model: "deepseek-v3"

task:
  max_iterations: 20
  timeout: 300  # seconds
  workspace: "./workspace"
```

## 7. 部署说明

### 7.1 依赖

- Go 1.21+
- Python 3.8+ (用于 CodeAgent 执行 Python 代码)
- pandas, openpyxl, matplotlib Python 包

### 7.2 启动方式

```bash
# 方式1: 直接运行
go run main.go

# 方式2: 编译运行
go build -o excelagent
./excelagent

# 方式3: Docker
docker build -t excelagent .
docker run -p 8080:8080 excelagent
```

## 8. 扩展点

- [ ] 支持异步任务模式（通过配置切换）
- [ ] 支持 Redis 存储任务状态
- [ ] 支持 WebSocket 实时推送进度
- [ ] 支持更多文件类型（PDF、Word）
- [ ] 支持流式输出（Server-Sent Events）
- [ ] 添加任务超时自动清理
- [ ] 添加任务优先级