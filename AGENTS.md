# AGENTS.md - Go MSA Project Guidelines

## MCP
- Use installed MCP servers for better answer:

## Build/Test Commands
- **Build**: `go build ./...`
- **Run**: `go run cmd/main.go`
- **Test all**: `go test ./...`
- **Test single**: `go test -run TestFunctionName ./path/to/package`
- **Test verbose**: `go test -v ./...`
- **Lint**: `golangci-lint run` (if installed)
- **Format**: `go fmt ./...`

## Code Style Guidelines

### Project Structure
- Follow clean architecture
- Use lowercase package names
- Group related functionality in packages

### Imports
- Standard library imports first
- Third-party imports second
- Local/project imports last
- Use blank lines to separate import groups

### Naming Conventions
- **Packages**: lowercase, descriptive (e.g., `domains`, `usecases`, `handlers`)
- **Functions**: PascalCase for exported, camelCase for unexported
- **Variables**: camelCase, descriptive names
- **Constants**: PascalCase for exported, camelCase for unexported

### Error Handling
- Return errors from functions, don't panic in business logic
- Use `log.Fatalf()` only for startup failures
- Handle errors at appropriate layers (handlers for HTTP responses)

### Types & Structs
- Define domain types in `internal/core/domain`
- Use pointer receivers for methods that modify structs
- Prefer structs over maps for structured data

### Dependencies
- Use dependency injection pattern
- Initialize dependencies in setup functions
- Close resources properly (defer statements)

### Logging
- Use `log` package for simple logging
- Include context in log messages
- Log startup/shutdown events</content>
<parameter name="filePath">/home/nokui/projects/go/go-msa/AGENTS.md