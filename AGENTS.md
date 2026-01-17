# Repository Guidelines

## Project Structure & Module Organization

- Root Go module: `go.mod`, `go.sum` (module `github.com/fatih/color`, Go `1.24.1`).
- Library source: `color.go`, `color_windows.go`, `doc.go`.
- Tests: `*_test.go` (e.g., `color_test.go`).
- Docs/notes: `README.md`, `TABWRITER_ALIGNMENT_ISSUE.md`, `LICENSE.md`.

## Build, Test, and Development Commands

- `go test ./...` — run the full test suite.
- `go test -run TestColor` — run a specific test.
- `go vet ./...` — static checks (recommended for changes).
- `go test -race ./...` — optional race detection for changes touching concurrency.

## Coding Style & Naming Conventions

- Follow standard Go formatting (`gofmt -w .`).
- Use idiomatic Go naming: PascalCase for exported identifiers, camelCase for unexported.
- Keep files small and focused; prefer adding platform-specific code to `*_windows.go`.

## Testing Guidelines

- Tests use Go’s `testing` package.
- Name tests `TestXxx` and focus on deterministic behavior (e.g., color escape sequences).
- Prefer table-driven tests when adding multiple cases.
- Run `go test ./...` before submitting changes.

## Commit & Pull Request Guidelines

- Recent commit messages are short, imperative, and descriptive (e.g., “Update CI and go deps”, “Bump golang.org/x/sys…”).
- Keep commits focused on a single change.
- PRs should describe the change, include rationale, and link relevant issues. Add before/after notes if output or API behavior changes.

## Configuration & Environment Notes

- Color output can be disabled via `NO_COLOR` or `color.NoColor`; include tests if behavior changes.
- Windows output uses `go-colorable`; validate platform-specific changes on Windows where possible.
