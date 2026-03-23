# einoSkill — Eino Skill 中间件演示

这是一个基于 **Gin** 的 HTTP 服务，演示了 [Eino](https://github.com/cloudwego/eino) 的 **Skill 中间件** 功能，灵感来源于 [eino-examples skills 示例](https://github.com/cloudwego/eino-examples)。

## 什么是 Skill？

Skill 是可复用的知识包——带有 YAML 前言的 markdown 文件——按需注入到 LLM 代理的上下文中。代理不需要将所有指令都塞进系统提示词，而是通过调用 `skill` 工具来加载当前任务所需的专业知识。

每个 Skill 都位于 `skills/` 目录下的独立子目录中：

```
skills/
├── go-best-practices/
│   └── SKILL.md     ← Go 惯用法、错误处理、命名规范、测试
└── code-review/
    └── SKILL.md     ← 代码审查清单（正确性、安全性、性能）
```

`SKILL.md` 文件格式如下：

```markdown
---
name: go-best-practices
description: Go 语言最佳实践和惯用模式。
---

# Go 最佳实践
...
```

## API

### `POST /chat`

通过 **Server-Sent Events (SSE)** 流式返回代理响应。

**请求体：**
```json
{"message": "Go 中应该如何处理错误？"}
```

**响应** — SSE 流：
```
event: message
data: 始终检查并返回错误 …

event: message
data: 使用 fmt.Errorf("context: %w", err) 来包装错误 …

data: [DONE]
```

### `GET /healthcheck`

服务运行时返回 `{"status":"ok"}`。

## 快速开始

```bash
export OPENAI_API_KEY="sk-..."
export OPENAI_MODEL="gpt-4o-mini"   # 可选，默认: gpt-4o-mini
# SKILLS_DIR 默认为源代码旁的 skills/ 子目录。
# 当从其他目录运行编译后的二进制文件时可覆盖此设置：
# export SKILLS_DIR="/path/to/skills"

cd bizCase/einoSkill
go run .
```

然后发送请求：

```bash
curl -N -X POST http://localhost:8080/chat \
  -H "Content-Type: application/json" \
  -d '{"message":"审查以下代码并列出问题：func div(a, b int) int { return a/b }"}'
```

## 架构

```
main.go          ← Gin 路由设置
init.go          ← OpenAI 模型 + skill 中间件 + 代理初始化
handler.go       ← SSE 流式聊天处理器
skill_backend.go ← 从本地 skills/ 目录读取 SKILL.md 文件
skills/          ← Skill 定义（SKILL.md 文件）
```

`localSkillBackend` 实现了 `skill.Backend` 接口，直接从磁盘读取 `SKILL.md` 文件，因此无需重新编译即可添加或修改 Skill。