# Code Style and Conventions

## Project Structure
- Follow clean architecture principles
- Use lowercase package names
- Group related functionality in packages

## Naming Conventions
- **Packages**: lowercase, descriptive (e.g., `domains`, `usecases`, `handlers`)
- **Functions**: PascalCase for exported, camelCase for unexported
- **Variables**: camelCase, descriptive names
- **Constants**: PascalCase for exported, camelCase for unexported

## Imports
- Standard library imports first
- Third-party imports second
- Local/project imports last
- Use blank lines to separate import groups

## Error Handling
- Return errors from functions, don't panic in business logic
- Use `log.Fatalf()` only for startup failures
- Handle errors at appropriate layers (handlers for HTTP responses)

## Types & Structs
- Define domain types in `internal/core/domain`
- Use pointer receivers for methods that modify structs
- Prefer structs over maps for structured data

## Dependencies
- Use dependency injection pattern
- Initialize dependencies in setup functions
- Close resources properly (defer statements)

## Logging
- Use `log` package for simple logging
- Include context in log messages
- Log startup/shutdown events