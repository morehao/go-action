---
name: go-best-practices
description: Go language best practices and idiomatic patterns, including error handling, concurrency, interfaces, and project structure guidelines.
---

# Go Best Practices

## Error Handling
- Always check and return errors; never ignore them silently.
- Use `fmt.Errorf("context: %w", err)` to wrap errors with context.
- Define sentinel errors with `var ErrXxx = errors.New(...)` for exported errors.
- Use `errors.Is` and `errors.As` to inspect wrapped errors.

## Naming Conventions
- Use short, concise names for local variables (e.g., `i`, `v`, `err`).
- Exported names should be meaningful and self-documenting.
- Interface names should end with `-er` when they wrap a single method (e.g., `Reader`, `Writer`).
- Avoid redundant package names (e.g., use `log.Info`, not `log.LogInfo`).

## Concurrency
- Prefer channels for communication between goroutines.
- Use `sync.WaitGroup` to wait for a collection of goroutines.
- Use `sync.Mutex` or `sync.RWMutex` to protect shared state.
- Always pass context as the first argument; cancel it to stop goroutines.

## Project Structure
- Keep `main.go` minimal; put business logic in packages.
- Group related functionality in sub-packages.
- Follow the standard layout: `cmd/`, `internal/`, `pkg/`.

## Testing
- Use table-driven tests for comprehensive coverage.
- Name tests as `TestXxx_conditionScenario_expectedResult`.
- Use `t.Helper()` in test helper functions.
- Mock external dependencies with interfaces.
