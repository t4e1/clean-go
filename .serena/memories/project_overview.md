# Project Overview

## Purpose
This is a Go-based Microservices Architecture (MSA) project implementing Clean Architecture. It demonstrates building scalable microservices using Go with hexagonal architecture principles.

## Tech Stack
- **Language**: Go 1.25.0
- **Web Framework**: Gin
- **Database**: PostgreSQL (via pgx/v5)
- **Authentication**: JWT
- **Other Libraries**: UUID, godotenv, testify for testing

## Design Patterns
- Clean Architecture (Hexagonal Architecture)
- Dependency Injection
- Ports and Adapters pattern

## Codebase Structure
- `cmd/`: Application entry points
- `configs/`: Configuration files
- `internal/core/domain/`: Domain entities and models
- `internal/core/usecases/`: Business logic
- `internal/ports/in/`: Input interfaces (APIs)
- `internal/ports/out/`: Output interfaces (repositories)
- `internal/adapters/in/`: Input adapters (HTTP handlers)
- `internal/adapters/out/`: Output adapters (database implementations)