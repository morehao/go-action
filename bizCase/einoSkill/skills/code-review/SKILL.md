---
name: code-review
description: Code review guidelines and checklist for reviewing Go code quality, security, performance, and maintainability.
---

# Code Review Guidelines

## Correctness
- Verify that the logic matches the intended behavior described in the task.
- Check all edge cases: empty input, nil pointers, boundary conditions.
- Ensure error paths are handled and do not cause panics or data loss.

## Security
- Never log sensitive data (passwords, tokens, PII).
- Validate and sanitize all user inputs before processing.
- Avoid SQL injection: use parameterized queries with your ORM/driver.
- Use `crypto/rand` instead of `math/rand` for security-sensitive randomness.

## Performance
- Avoid unnecessary allocations in hot paths; reuse buffers with `sync.Pool`.
- Prefer `strings.Builder` over `+` concatenation in loops.
- Use appropriate data structures (maps vs slices) based on access patterns.
- Defer heavy computations or I/O to background goroutines when latency matters.

## Readability
- Each function should do one thing; keep functions short and focused.
- Add comments to explain *why*, not *what* the code does.
- Avoid deeply nested conditionals; use early returns (guard clauses).
- Eliminate dead code and unused imports.

## Testing
- New features must have corresponding unit tests.
- Refactored code must not reduce existing test coverage.
- Tests should be deterministic and not rely on external services.
