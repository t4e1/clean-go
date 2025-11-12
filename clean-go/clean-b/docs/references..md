## 1. hexagonal architecture
- https://medium.com/@matiasvarela/hexagonal-architecture-in-go-cfd4e436faa3

```
internal/
├─ core/          ← Domain logic
│   ├─ domain/    ← domain model(entity)
│   └─ usecase/   ← 실제 비즈니스 로직 위치(service)
├─ adapters/
│   ├─ in/        ← HTTP, gRPC, CLI 등 실제 구현체 
│   └─ out/       ← DB, MQ 등 실제 구현체
├─ ports/
│   ├─ in/        ← HTTP, gRPC, CLI 등의 추상화(인터페이스)
│   └─ out/       ← DB, MQ 등의 추상화(인터페이스)
└─ configs/       ← 공통 유틸, 설정, 로깅

```