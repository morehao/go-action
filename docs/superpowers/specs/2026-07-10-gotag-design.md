# gotag 设计文档：Go Tag 实现开源/企业版代码隔离

## 概述

为 Go 项目设计一套完整的方案，用于在同一个代码仓库中通过 Go Build Tag 区分**开源版**和**企业版**，并实现 GitLab（完整代码）到 GitHub（仅开源代码）的单向代码同步。

## 核心策略

- **只定义一个 tag `enterprise`**：企业版构建时带 `-tags enterprise`，开源版是默认（无 tag）
- **`!enterprise` 等价于开源版**：一套 tag 表达两种语义，避免双 tag 管理的复杂性

## 目录结构（demo）

```
bizcase/gotag/
├── go.mod                          # module: github.com/morehao/go-action/bizcase/gotag
├── main.go                         # 入口，选择 Platform 接口实现
├── Makefile                        # 构建命令封装
├── sync-to-github.sh               # 同步脚本：gitlab → github
│
├── platform/                       # 核心接口定义（通用，无 tag）
│   └── platform.go                 # Platform 接口 + 通用类型
│
├── impl/                           # 实现层（按 tag 区分）
│   ├── oss.go                      //go:build !enterprise    本地文件存储
│   ├── oss_enterprise.go           //go:build enterprise      阿里云 OSS
│   ├── audit.go                    //go:build enterprise      企业版独有：审计日志
│   ├── audit_stub.go               //go:build !enterprise    开源版：空实现
│   ├── auth.go                     //go:build !enterprise    JWT 简单鉴权
│   └── auth_enterprise.go          //go:build enterprise      RBAC 鉴权
│
├── enterprise/                     # 企业版独占目录（整个目录）
│   └── report.go                   //go:build enterprise      报表服务
│
└── cmd/
    └── build-check/
        └── main.go                 # CI 构建验证
```

## 文件命名约定

| 文件后缀 | Build Tag | 含义 |
|----------|-----------|------|
| `*.go` | 无 | 通用代码，两版都编译 |
| `*_enterprise.go` | `//go:build enterprise` | 企业版专属实现 |
| `*_opensource.go` | `//go:build !enterprise` | 开源版专属实现（同步时重命名去掉后缀） |
| `*_stub.go` | `//go:build !enterprise` | 开源版空实现/桩 |

## 差异模式

### 模式 1：功能有无（企业版独有）

企业版有审计日志，开源版为空实现：

```go
// impl/audit.go  //go:build enterprise
type AuditLogger struct{ client *elastic.Client }
func (a *AuditLogger) Log(action, user string) { /* 写入 ES */ }

// impl/audit_stub.go  //go:build !enterprise
type AuditLogger struct{}
func (a *AuditLogger) Log(action, user string) { /* no-op */ }
```

### 模式 2：实现方式不同（对偶文件）

文件存储：企业版用阿里云 OSS，开源版用本地文件：

```go
// impl/oss.go  //go:build !enterprise
func NewOSS() OSS { return &LocalStorage{} }

// impl/oss_enterprise.go  //go:build enterprise
func NewOSS() OSS { return &AliyunOSS{endpoint: cfg.OSSEndpoint} }
```

### 模式 3：整个目录独占

`enterprise/report.go` 整个文件加 tag，同步时删除整个目录。

## 构建命令

```makefile
build-enterprise:
    go build -tags enterprise -o bin/server-enterprise .

build-opensource:
    go build -o bin/server-opensource .

build-all: build-opensource build-enterprise
```

CI 中两个构建都必须通过。

## 代码同步机制

### 同步流

```
GitLab（完整代码）── sync-to-github.sh ──▶ GitHub（仅开源）
```

单向推送，GitHub 的修改通过 PR 合并回 GitLab。

### 同步脚本逻辑

1. 在 GitLab 仓库创建临时分支 `sync-to-github/YYYYMMDD`
2. 删除所有匹配 `//go:build enterprise` 的文件（含 `_enterprise.go` 后缀文件）
3. 删除 `enterprise/` 等独占目录
4. 将 `*_opensource.go` 重命名去掉 `_opensource` 后缀
5. 清理 `go.mod` 中企业版私有依赖（匹配 `enterprise-only` 注释行）
6. 提交、打 tag
7. `git push` 到 GitHub 远程

### 双保险

- **源头控制**：同步脚本自动过滤企业版代码
- **接收方校验**：GitHub CI 中搜索是否存在 `//go:build enterprise` 或 `_enterprise.go`，存在则 fail

## go.mod 依赖管理

企业版私有依赖在 `go.mod` 中加 `enterprise-only` 注释标记，同步脚本自动移除：

```
require (
    github.com/gin-gonic/gin v1.12.0              // common
    github.com/aliyun/aliyun-oss-go-sdk v2.2.0    // enterprise-only
)
```

## 验证规则

- `go build -tags enterprise ./...` — 企业版必须通过
- `go build -o /dev/null ./...` — 开源版（默认无 tag）必须通过
- GitHub CI 禁止包含 `//go:build enterprise` 的代码
