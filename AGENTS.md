# AGENTS.md

## Build
```bash
# Build the binary
make build
```

## Lint
```bash
# Run golangci‑lint for static analysis
# (requires golangci‑lint installed)
make lint
```

## Test
```bash
# Run all tests
make test

# Run a single test by name
make test RUN=TestName
```

## Code style guidelines
- Run `go fmt ./...` and `go vet ./...` before committing.
- Imports should be grouped: standard library, then third‑party, then local packages.
- Use `goimports` to automatically sort and tidy imports.
- Exported identifiers: PascalCase.
- Unexported identifiers: lowerCamelCase.
- Avoid underscore prefixes; use `err` for errors.
- Prefer context over timeouts where possible.
- Keep functions small (≤ 15 lines) and document with comments.
- Use `errors.Is/As` for error handling.
- Write tests with clear Setup/Act/Assert structure.
- All test functions start with `Test` and are exported.
- Use table driven tests for multiple cases.
- Keep README updated with build/test instructions.
