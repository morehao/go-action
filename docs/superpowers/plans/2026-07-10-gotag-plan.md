# Go Tag 开源/企业版隔离 Implementation Plan

> **For agentic workers:** 此方案已实现完毕，本文档为回顾性总结，记录实际项目迁移的实施步骤。

**目标：** 在 Go 项目中通过 build tag 实现开源版和企业版的代码隔离，并建立 GitLab→GitHub 的单向同步机制。

**架构：** 单 tag（`enterprise`）策略，接口抽象 + 对偶文件 + stub 模式的组合方案。

**Tech Stack:** Go 1.26+, Makefile, GitLab CI, GitHub Actions

## 全局约束

- 只定义 `enterprise` 一个 build tag，开源版为默认无 tag
- tag 语法使用 `//go:build`（Go 1.17+）
- 同步只能从 GitLab→GitHub 单向，GitHub 不反向同步
- 企业版私有依赖在 go.mod 中用 `// enterprise-only` 注释标记

---

### Task 1: 项目目录结构搭建

**文件：**
- 创建: `bizcase/gotag/` 整个目录结构
- 创建: `bizcase/gotag/go.mod`
- 创建: `bizcase/gotag/platform/platform.go` — 核心接口定义

**接口：**
- 产出: `platform.OSS` 接口 — `Upload(bucket, key string, data []byte) error` + `Download(bucket, key string) ([]byte, error)`
- 产出: `platform.Auth` 接口 — `Login(username, password string) (string, error)` + `Validate(token string) (string, error)`
- 产出: `platform.AuditLog` 接口 — `Log(action, user string)`
- 产出: `platform.Report` 接口 — `GenerateReport(title string) (string, error)`
- 产出: `platform.Platform` 结构体 — 持有以上四个接口，提供 `Run()` 方法
- 产出: `platform.NewPlatform(oss, auth, audit, report)` 工厂函数

- [ ] **Step 1: 创建目录结构**

```bash
mkdir -p bizcase/gotag/{platform,impl,enterprise,cmd/build-check}
```

- [ ] **Step 2: 创建 go.mod**

```
module github.com/morehao/go-action/bizcase/gotag

go 1.26.1
```

- [ ] **Step 3: 创建 platform 接口定义**

`bizcase/gotag/platform/platform.go` — 定义 OSS、Auth、AuditLog、Report 四个接口和 Platform 聚合结构体。

- [ ] **Step 4: 验证编译**

```bash
go build -o /dev/null .
```

---

### Task 2: 开源版实现层（`!enterprise` tag）

**文件：**
- 创建: `bizcase/gotag/impl/oss.go` — `//go:build !enterprise`，LocalStorage 实现
- 创建: `bizcase/gotag/impl/auth.go` — `//go:build !enterprise`，JWTAuth 实现
- 创建: `bizcase/gotag/impl/audit_stub.go` — `//go:build !enterprise`，StubAuditLog 空实现
- 创建: `bizcase/gotag/enterprise/report_stub.go` — `//go:build !enterprise`，ReportService 返回禁用错误

**接口：**
- 消费: `platform.OSS`, `platform.Auth`, `platform.AuditLog`, `platform.Report`
- 产出: `impl.NewOSS()` → `*LocalStorage`
- 产出: `impl.NewAuth()` → `*JWTAuth`
- 产出: `impl.NewAuditLog()` → `*StubAuditLog`
- 产出: `enterprise.NewReportService()` → `*enterprise.ReportService`

- [ ] **Step 1: 实现 LocalStorage**

`impl/oss.go` 使用本地文件系统存储，rootDir 为 `/tmp/gotag-storage`。`Upload` 创建目录并写文件，`Download` 读文件。

- [ ] **Step 2: 实现 JWTAuth**

`impl/auth.go` 简单 JWT 风格鉴权，`Login` 校验用户名密码，`Validate` 检查 token 长度。

- [ ] **Step 3: 实现 AuditLog stub**

`impl/audit_stub.go` — `Log` 方法为空操作。

- [ ] **Step 4: 实现 Report stub**

`enterprise/report_stub.go` — `GenerateReport` 返回错误 "report is only available in enterprise edition"。

- [ ] **Step 5: 验证编译**

```bash
go build -o /dev/null .
go test ./...
```

---

### Task 3: 企业版实现层（`enterprise` tag）

**文件：**
- 创建: `bizcase/gotag/impl/oss_enterprise.go` — `//go:build enterprise`，AliyunOSS 实现
- 创建: `bizcase/gotag/impl/auth_enterprise.go` — `//go:build enterprise`，RBACAuth 实现
- 创建: `bizcase/gotag/impl/audit.go` — `//go:build enterprise`，ESAuditLog 实现
- 创建: `bizcase/gotag/enterprise/report.go` — `//go:build enterprise`，ReportService 真实实现

**接口：**
- 消费: `platform.OSS`, `platform.Auth`, `platform.AuditLog`, `platform.Report`
- 产出: `impl.NewOSS()` → `*AliyunOSS`
- 产出: `impl.NewAuth()` → `*RBACAuth`
- 产出: `impl.NewAuditLog()` → `*ESAuditLog`
- 产出: `enterprise.NewReportService()` → `*enterprise.ReportService`

- [ ] **Step 1: 实现 AliyunOSS**

`impl/oss_enterprise.go` — 模拟阿里云 OSS 上传下载，打印日志模拟实际调用。

- [ ] **Step 2: 实现 RBACAuth**

`impl/auth_enterprise.go` — 模拟 LDAP 认证，token 包含角色信息。

- [ ] **Step 3: 实现 ESAuditLog**

`impl/audit.go` — 模拟 Elasticsearch 审计日志写入，记录时间、操作、用户。

- [ ] **Step 4: 实现 Report 真实实现**

`enterprise/report.go` — 返回包含 "Enterprise Report" 前缀的报表内容。

- [ ] **Step 5: 验证编译**

```bash
go build -tags enterprise -o /dev/null .
go test -tags enterprise ./...
```

---

### Task 4: 入口和构建体系

**文件：**
- 创建: `bizcase/gotag/main.go` — 入口，组装所有组件并运行
- 创建: `bizcase/gotag/Makefile` — 构建命令封装
- 创建: `bizcase/gotag/cmd/build-check/main.go` — CI 构建验证

- [ ] **Step 1: 创建 main.go**

聚合 impl 和 enterprise 的工厂函数，创建 Platform 实例并调用 Run()，然后依次测试 Login、Upload、Audit.Log、Report。

- [ ] **Step 2: 创建 Makefile**

```
build-opensource:  go build -o bin/server-opensource .
build-enterprise:  go build -tags enterprise -o bin/server-enterprise .
build-all:         build-opensource + build-enterprise
run-opensource:    build-opensource + run
run-enterprise:    build-enterprise + run
test:              test both editions
verify:            grep 检查没有 enterprise tag 泄漏
```

- [ ] **Step 3: 创建 CI 构建验证脚本**

`cmd/build-check/main.go` — 分别执行 `go build` 和 `go build -tags enterprise`，两个都通过才算成功。

- [ ] **Step 4: 验证两版运行**

```bash
# 开源版
go run .  # 输出: LocalStorage, JWTAuth, Stub, Report error
# 企业版
go run -tags enterprise .  # 输出: AliyunOSS, RBACAuth, ESAudit, Report success
```

---

### Task 5: 测试

**文件：**
- 创建: `bizcase/gotag/impl/oss_test.go` — `//go:build !enterprise`，开源版测试
- 创建: `bizcase/gotag/impl/oss_enterprise_test.go` — `//go:build enterprise`，企业版测试

- [ ] **Step 1: 编写开源版测试**

```go
//go:build !enterprise
package impl
func TestLocalStorage_UploadAndDownload(t *testing.T) { ... }
func TestJWTAuth_Login(t *testing.T) { ... }
func TestJWTAuth_Login_Invalid(t *testing.T) { ... }
```

- [ ] **Step 2: 编写企业版测试**

```go
//go:build enterprise
package impl
func TestAliyunOSS_Upload(t *testing.T) { ... }
func TestRBACAuth_Login(t *testing.T) { ... }
```

- [ ] **Step 3: 运行测试**

```bash
go test ./...                # 开源版测试通过
go test -tags enterprise ./...  # 企业版测试通过
```

---

### Task 6: 同步脚本

**文件：**
- 创建: `bizcase/gotag/sync-to-github.sh` — GitLab→GitHub 单向同步

- [ ] **Step 1: 创建同步脚本**

核心逻辑：
1. 创建临时分支 `sync-to-github/YYYYMMDD-HHMMSS`
2. 删除所有含 `//go:build enterprise` 的 .go 文件
3. 将 `*_opensource.go` 重命名为去掉 `_opensource` 后缀
4. 清理 go.mod 中 `enterprise-only` 标记的依赖行
5. `go mod tidy` 整理依赖
6. 提交、push 到 GitHub 远程

- [ ] **Step 2: 添加执行权限**

```bash
chmod +x sync-to-github.sh
```

---

### Task 7: 设计文档和提交

- [ ] **Step 1: 编写设计文档**

`docs/superpowers/specs/2026-07-10-gotag-design.md`

- [ ] **Step 2: 提交所有代码**

```bash
git add bizcase/gotag/ docs/superpowers/specs/2026-07-10-gotag-design.md
git commit -m "feat: add gotag demo - Go build tag for open-source/enterprise edition isolation"
```

---

## 实际项目迁移指南

### 步骤 1：在真实项目中建立目录结构

```
your-project/
├── internal/
│   ├── platform/       # 接口定义（通用）
│   ├── impl/           # 实现层（按 tag 区分）
│   └── enterprise/     # 企业版独占目录
├── Makefile            # 构建命令
├── sync-to-github.sh   # 同步脚本
└── .github/workflows/  # CI 检查
```

### 步骤 2：识别现有代码中的差异点

| 差异类型 | 迁移方式 | 示例 |
|----------|----------|------|
| 功能实现不同 | 提取接口，创建对偶文件 | `oss.go` + `oss_enterprise.go` |
| 企业版独有功能 | 企业版真实实现 + 开源版 stub | `audit.go` + `audit_stub.go` |
| 企业版独占模块 | 整个目录加 tag | `enterprise/` 目录 |

### 步骤 3：配置 CI

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

### 步骤 4：建立同步流水线

在 GitLab 配置定时任务（如每日凌晨）执行 `sync-to-github.sh`，自动同步开源代码到 GitHub。