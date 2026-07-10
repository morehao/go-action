# gotag — Go Build Tag 开源版/企业版代码隔离 Demo

## 概述

本 Demo 演示如何通过 Go Build Tag（`//go:build`）在**同一个代码仓库**中实现**开源版**和**企业版**的代码隔离，并建立从 GitLab（完整代码）到 GitHub（仅开源代码）的单向同步机制。

模拟场景：一个 SaaS 平台，包含 OSS 存储、用户鉴权、审计日志、报表服务四个功能，企业版有更强大的实现和额外功能。

## 核心策略

**只定义一个 tag `enterprise`：企业版构建时带 `-tags enterprise`，开源版是默认（无 tag）。**

| 文件后缀 | Build Tag | 含义 |
|----------|-----------|------|
| `*.go` | 无 | 通用代码，两版都编译 |
| `*_enterprise.go` | `//go:build enterprise` | 企业版专属实现 |
| `*_opensource.go` | `//go:build !enterprise` | 开源版专属实现（同步时去掉 `_opensource` 后缀） |
| `*_stub.go` | `//go:build !enterprise` | 开源版空实现/桩 |

## 三种差异模式

### 模式 1：功能有无（企业版独有功能）

企业版有审计日志（写入 Elasticsearch），开源版为空操作：

```go
// impl/audit.go          //go:build enterprise
type AuditLogger struct{ client *elastic.Client }
func (a *AuditLogger) Log(action, user string) { /* 写入 ES */ }

// impl/audit_stub.go     //go:build !enterprise
type AuditLogger struct{}
func (a *AuditLogger) Log(action, user string) { /* no-op */ }
```

### 模式 2：实现方式不同（对偶文件）

文件存储：企业版用阿里云 OSS，开源版用本地文件：

```go
// impl/oss.go            //go:build !enterprise
func NewOSS() OSS { return &LocalStorage{} }

// impl/oss_enterprise.go //go:build enterprise
func NewOSS() OSS { return &AliyunOSS{endpoint: "..."} }
```

用户鉴权：企业版用 LDAP RBAC，开源版用简单 JWT：

```go
// impl/auth.go           //go:build !enterprise
func NewAuth() Auth { return &JWTAuth{} }

// impl/auth_enterprise.go //go:build enterprise
func NewAuth() Auth { return &RBACAuth{ldapURL: "..."} }
```

### 模式 3：整个目录独占

报表服务整个目录只属于企业版，开源版返回"仅企业版可用"错误：

```go
// enterprise/report.go       //go:build enterprise
func (r *ReportService) GenerateReport(title string) (string, error) {
    return fmt.Sprintf("[Enterprise Report] %s", title), nil
}

// enterprise/report_stub.go  //go:build !enterprise
func (r *ReportService) GenerateReport(title string) (string, error) {
    return "", fmt.Errorf("report is only available in enterprise edition")
}
```

## 快速开始

```bash
# 运行开源版
make run-opensource

# 运行企业版
make run-enterprise

# 测试两版
make test

# 两版同时构建
make build-all

# 检查企业版代码泄漏（用于 GitHub CI）
make verify
```

### 开源版输出

```
=== Platform Running ===
OSS: &{/tmp/gotag-storage}
Auth: &{opensource-secret}
Audit: &{}
Report: &{}

=== Testing Features ===
Login success, token: token_admin
Upload success
Report failed: report is only available in enterprise edition
```

### 企业版输出

```
=== Platform Running ===
OSS: &{https://oss-cn-hangzhou.aliyuncs.com enterprise-bucket}
Auth: &{ldap://enterprise-ldap.internal:389}
Audit: &{http://elasticsearch.internal:9200}
Report: &{}

=== Testing Features ===
[RBAC] Authenticating admin via ldap://enterprise-ldap.internal:389
Login success, token: enterprise_token_admin_role_admin
[AliyunOSS] Uploading my-bucket/hello.txt ...
Upload success
[Audit][ES] time=... action=file.upload user=admin ...
[Enterprise Report] Monthly Summary - generated with advanced BI engine
```

## 目录结构

```
bizcase/gotag/
├── go.mod                          # module: github.com/morehao/go-action/bizcase/gotag
├── main.go                         # 入口，组装各组件并运行演示
├── Makefile                        # 构建命令封装
├── sync-to-github.sh               # GitLab → GitHub 单向同步脚本
│
├── platform/                       # 核心接口定义（通用，无 tag）
│   └── platform.go                 # OSS / Auth / AuditLog / Report 接口 + Platform 聚合
│
├── impl/                           # 接口实现层（按 tag 区分）
│   ├── oss.go                      //go:build !enterprise    本地文件存储
│   ├── oss_enterprise.go           //go:build enterprise      阿里云 OSS
│   ├── auth.go                     //go:build !enterprise    JWT 简单鉴权
│   ├── auth_enterprise.go          //go:build enterprise      LDAP RBAC 鉴权
│   ├── audit.go                    //go:build enterprise      Elasticsearch 审计日志
│   ├── audit_stub.go               //go:build !enterprise    审计日志空实现
│   ├── oss_test.go                 //go:build !enterprise    开源版测试
│   └── oss_enterprise_test.go      //go:build enterprise      企业版测试
│
├── enterprise/                     # 企业版独占目录
│   ├── report.go                   //go:build enterprise      报表服务真实实现
│   └── report_stub.go              //go:build !enterprise    报表服务桩（返回禁用错误）
│
└── cmd/
    └── build-check/
        └── main.go                 # CI 构建验证：同时编译两版
```

## 代码同步机制

### 同步流

```
GitLab（完整代码，含企业版）
    │
    │  sync-to-github.sh
    ▼
GitHub（仅开源代码）
```

同步是**单向**的。GitHub 上的修改需通过 PR 合并回 GitLab。

### 同步脚本逻辑

1. 创建临时分支 `sync-to-github/YYYYMMDD-HHMMSS`
2. 删除所有含 `//go:build enterprise` 的 `.go` 文件（含 `_enterprise.go` 后缀）
3. 删除 `enterprise/` 等独占目录
4. 将 `*_opensource.go` 重命名去掉 `_opensource` 后缀
5. 清理 `go.mod` 中 `enterprise-only` 标记的企业版私有依赖
6. 提交、push 到 GitHub

### 双保险

- **源头控制**：同步脚本自动过滤企业版代码
- **接收方校验**：GitHub CI 中 `make verify` 搜索是否存在 `//go:build enterprise`，存在则 fail

### 双仓库模拟

```bash
# 在 GitLab 模拟仓库中执行，推送到 GitHub 模拟仓库
cd bizcase/gotag
git init
git add .
git commit -m "init: full code with enterprise edition"

# 同步到 GitHub 模拟仓库
./sync-to-github.sh git@github.com:yourname/gotag-opensource.git
```

## 实际项目迁移指南

### 第一步：建立目录结构

```
your-project/
├── internal/
│   ├── platform/       # 接口定义（通用）
│   ├── impl/           # 实现层（按 tag 区分）
│   └── enterprise/     # 企业版独占目录
├── Makefile
├── sync-to-github.sh
└── .github/workflows/
```

### 第二步：识别现有代码中的差异点

| 差异类型 | 迁移方式 | 本 Demo 示例 |
|----------|----------|-------------|
| 功能实现不同 | 提取接口，创建对偶文件 | `oss.go` + `oss_enterprise.go` |
| 企业版独有功能 | 企业版实现 + 开源版 stub | `audit.go` + `audit_stub.go` |
| 企业版独占模块 | 整个目录所有文件加 tag | `enterprise/` 目录 |

### 第三步：配置 CI

**GitLab CI：**
```yaml
build-enterprise:
  script: go build -tags enterprise ./...

build-opensource:
  script: go build ./...
```

**GitHub Actions：**
```yaml
- name: Check for enterprise tag leaks
  run: '! grep -r "//go:build enterprise" --include="*.go" .'
```

### 第四步：建立同步流水线

在 GitLab 配置定时任务（如每日凌晨）执行 `sync-to-github.sh`，自动同步开源代码到 GitHub。
