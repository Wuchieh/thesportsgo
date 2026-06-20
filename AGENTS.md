# AGENTS.md — thesportsgo

**Generated:** 2026-06-20 | **Commit:** e2bec4a | **Branch:** master

Go client library for the TheSports.com API v1 (`github.com/wuchieh/thesportsgo`). Zero external dependencies. Flat single-package design.

## Commands

```bash
# Lint (uses gofumpt, not gofmt)
golangci-lint run ./...

# Tests — none exist yet, so this will report "no test files" (expected)
go test ./...
```

## Architecture

- **Flat package** — everything lives in root package `thesportsgo`. No subpackages.
- **`Client`** (`client.go`) — holds `user`/`secret` auth and `http.Client`. Created via `NewClient(user, secret, opts...)`. Base URL defaults to `https://api.thesports.com/v1`.
- **Functional options** — `WithHTTPClient`, `WithBaseURL`.
- **Generic API helper** — `secretGet[T Response[E], E any]` in `utils.go` is the single call path for all API endpoints. It attaches auth params, sends GET, unmarshals to `MetaResponse[E]`, checks errors, and returns `*Response[E]`.
- **Response types** — `Response[T]` (public, `code` + `results`) and `MetaResponse[T]` (internal, also `err`/`msg`, implements `error`). `MetaResponse.UnmarshalJSON` captures the raw response body for debugging.
- **Query structs** — `toQuery(any)` converts a struct to `url.Values` by JSON round-trip + reflection. Pointer fields with `omitempty` json tags are skipped when nil.

## WHERE TO LOOK

| Task | Location | Notes |
|------|----------|-------|
| Add new sport API | `football.go` + `football_url.go` + `football_response.go` | Follow naming: `{Sport}{Endpoint}`, `{sport}{Endpoint}Path`, `{Sport}{Endpoint}ResponseData` |
| Core client changes | `client.go` | `NewClient`, `Get`, `SecretGet`, options |
| Response handling | `response.go` | `Response[T]`, `MetaResponse[T]`, custom `UnmarshalJSON` |
| Generic helpers | `utils.go` | `secretGet` (single call path), `toQuery` (struct→query) |
| Integration test | `_test/main.go` | Manual script, gitignored, `go run ./_test/` |

## CODE MAP

| Symbol | Type | Location | Role |
|--------|------|----------|------|
| `NewClient` | Func | `client.go:89` | Entry point — creates API client |
| `SecretGet` | Method | `client.go:35` | Attach auth + GET |
| `secretGet[T,E]` | Func | `utils.go:86` | **Single call path for all endpoints** |
| `toQuery` | Func | `utils.go:49` | Struct → url.Values (JSON round-trip) |
| `Response[T]` | Struct | `response.go:10` | Public result wrapper |
| `MetaResponse[T]` | Struct | `response.go:22` | Internal wrapper; implements `error` |
| `FootballCategory` | Method | `football.go:10` | Example endpoint pattern |
| `footballCategoryPath` | Const | `football_url.go:4` | Endpoint URL constants |

## ANTI-PATTERNS (THIS PROJECT)

- **NEVER** use `as any`, `@ts-ignore`, or type assertions to bypass type errors (Go: no `interface{}` casts to escape generics).
- **NEVER** introduce external dependencies — this module is intentionally dependency-free.
- **NEVER** break the flat package structure — no subpackages.
- **DO NOT** use `gofmt` — formatter is `gofumpt` (see golangci.yml).
- **DO NOT** remove `omitempty` from pointer fields in query structs — this is how `toQuery` skips unset fields.
- **DO NOT** pass `nil` context to API methods — though `handleContext` handles it, it's non-idiomatic.

## UNIQUE STYLES

- Comments: traditional Chinese on exported identifiers.
- Response flow: `MetaResponse` (raw) → error check → `Response` (public).
- Type aliases for endpoint responses: `FootballCategoryResponse = Response[[]FootballCategoryResponseData]`.

## NOTES

- `golangci-lint` uses `tests: false` — test files are NOT linted. If `_test.go` files are added, update to `true` or remove the line.
- `must[T]` in `utils.go:42` is dead code (unused). OK to remove.
- `handlePath` in `client.go:57` has a `staticcheck S1017` warning — should use `strings.TrimPrefix`.
- `FootballCompetition` returns `*FootballCountryResponse` — may be a copy-paste bug; verify if competition data matches country response shape.
- `toQuery` panics on marshal errors (acceptable for internal use, but library callers may expect error returns).
- `_test/main.go` is gitignored — `go test ./...` reports "no test files" (expected).

## Conventions

- **Formatter**: `gofumpt` (not plain `gofmt`).
- **Linter config**: `golangci.yml` enables `revive`, `staticcheck`, `gocyclo` (min 20), `gofumpt`, etc. `go test` is excluded from linting (`tests: false`).
- **Go version**: 1.25.5, module has zero external dependencies.
