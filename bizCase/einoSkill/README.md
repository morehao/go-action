# einoSkill — Eino Skill Middleware Demo

A **Gin**-based HTTP service that demonstrates [Eino](https://github.com/cloudwego/eino)'s **Skill Middleware** feature, inspired by the [eino-examples skills examples](https://github.com/cloudwego/eino-examples).

## What Are Skills?

Skills are reusable knowledge packages—markdown files with YAML frontmatter—that are injected into an LLM agent's context on demand. Instead of stuffing every instruction into the system prompt, the agent calls a `skill` tool to load only the expertise it needs for the current task.

Each skill lives in its own sub-directory under `skills/`:

```
skills/
├── go-best-practices/
│   └── SKILL.md     ← Go idioms, error handling, naming, testing
└── code-review/
    └── SKILL.md     ← Review checklist (correctness, security, perf)
```

A `SKILL.md` file looks like this:

```markdown
---
name: go-best-practices
description: Go language best practices and idiomatic patterns.
---

# Go Best Practices
...
```

## API

### `POST /chat`

Stream a response from the agent via **Server-Sent Events (SSE)**.

**Request body:**
```json
{"message": "How should I handle errors in Go?"}
```

**Response** — SSE stream:
```
event: message
data: Always check and return errors …

event: message
data: Use fmt.Errorf("context: %w", err) to wrap them …

data: [DONE]
```

### `GET /healthcheck`

Returns `{"status":"ok"}` when the server is up.

## Quick Start

```bash
export OPENAI_API_KEY="sk-..."
export OPENAI_MODEL="gpt-4o-mini"   # optional, default: gpt-4o-mini
# SKILLS_DIR defaults to the skills/ sub-directory next to the source files.
# Override it when running a compiled binary from another directory:
# export SKILLS_DIR="/path/to/skills"

cd bizCase/einoSkill
go run .
```

Then send a request:

```bash
curl -N -X POST http://localhost:8080/chat \
  -H "Content-Type: application/json" \
  -d '{"message":"Review the following code and list any issues: func div(a, b int) int { return a/b }"}'
```

## Architecture

```
main.go          ← Gin router setup
init.go          ← OpenAI model + skill middleware + agent initialisation
handler.go       ← SSE streaming chat handler
skill_backend.go ← Reads SKILL.md files from the local skills/ directory
skills/          ← Skill definitions (SKILL.md files)
```

The `localSkillBackend` implements `skill.Backend` and reads `SKILL.md` files directly from disk, so you can add or modify skills without recompiling.
