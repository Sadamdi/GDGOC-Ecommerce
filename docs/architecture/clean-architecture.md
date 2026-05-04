# Clean Architecture — E-Commerce Backend

> **Prinsip Utama**: Dependency hanya mengarah ke dalam. Layer luar bergantung pada layer dalam, TIDAK PERNAH sebaliknya.

---

## 🎯 Arsitektur Overview

```mermaid
graph TB
    subgraph "Layer 4: Infrastructure"
        H[HTTP Handler]
        MW[Middleware]
        SW[Swagger UI]
    end
    
    subgraph "Layer 3: Interface Adapters"
        R[Repository Implementation]
        EXT[External Services]
    end
    
    subgraph "Layer 2: Application"
        UC[Use Cases]
    end
    
    subgraph "Layer 1: Domain - CORE"
        E[Entities]
        RI[Repository Interfaces]
        UCI[UseCase Interfaces]
        DE[Domain Errors]
        DV[Domain Validation]
    end
    
    H --> UC
    MW --> UC
    UC --> RI
    UC --> E
    R --> RI
    R --> E
    EXT --> RI
    
    style E fill:#4caf50,color:#fff
    style RI fill:#4caf50,color:#fff
    style UCI fill:#4caf50,color:#fff
    style DE fill:#4caf50,color:#fff
    style DV fill:#4caf50,color:#fff
    style UC fill:#2196f3,color:#fff
    style R fill:#ff9800,color:#fff
    style EXT fill:#ff9800,color:#fff
    style H fill:#f44336,color:#fff
    style MW fill:#f44336,color:#fff
    style SW fill:#f44336,color:#fff
```

---

## 📐 Layer Rules

| Layer | Boleh Import | Tidak Boleh Import |
|-------|-------------|-------------------|
| **Domain** | Standard library saja | UseCase, Repository impl, Handler, third-party DB |
| **UseCase** | Domain | Repository impl, Handler, HTTP library |
| **Repository** | Domain, DB driver | UseCase, Handler |
| **Handler** | Domain, UseCase interface | Repository impl, DB driver |

---

## 🔄 Request Flow

```mermaid
sequenceDiagram
    participant C as Client
    participant MW as Middleware
    participant H as Handler
    participant UC as UseCase
    participant R as Repository
    participant DB as MongoDB
    
    C->>MW: HTTP Request
    MW->>MW: Auth, CORS, Logging
    MW->>H: Validated Request
    H->>H: Decode & Validate Input
    H->>UC: Call UseCase Method
    UC->>UC: Business Logic
    UC->>R: Call Repository (via Interface)
    R->>DB: Database Query
    DB-->>R: Result
    R-->>UC: Domain Entity
    UC-->>H: Response DTO
    H-->>C: JSON Response
```

---

## 🏛️ Dependency Injection Flow

```mermaid
flowchart TD
    A[main.go] --> B[Initialize Config]
    B --> C[Connect MongoDB]
    C --> D[Create Repository Instances]
    D --> E[Create UseCase Instances]
    E --> F[Create Handler Instances]
    F --> G[Setup Router & Middleware]
    G --> H[Start HTTP Server]
    
    D -->|inject repo| E
    E -->|inject usecase| F
```

```go
// cmd/api/main.go — Dependency Injection
func main() {
    cfg := config.Load()
    db := mongodb.Connect(cfg.MongoURI)
    
    // Repository layer
    userRepo := mongoRepo.NewUserRepository(db)
    productRepo := mongoRepo.NewProductRepository(db)
    
    // UseCase layer — inject repository interfaces
    userUC := usecase.NewUserUseCase(userRepo, hasher, tokenGen)
    productUC := usecase.NewProductUseCase(productRepo)
    
    // Handler layer — inject usecase interfaces
    userHandler := handler.NewUserHandler(userUC)
    productHandler := handler.NewProductHandler(productUC)
    
    // Router
    router := http.NewServeMux()
    registerRoutes(router, userHandler, productHandler)
    
    http.ListenAndServe(":8080", router)
}
```

---

## 📏 Layer Violation Detector

Gunakan rules ini untuk mendeteksi pelanggaran arsitektur:

| Violation | Contoh | Perbaikan |
|-----------|--------|-----------|
| Handler import DB driver | `import "go.mongodb.org/..."` di handler | Gunakan UseCase interface |
| UseCase import HTTP | `import "net/http"` di usecase | Gunakan domain DTOs |
| Domain import third-party | `import "go.mongodb.org/..."` di domain | Hanya standard library |
| Repository berisi business logic | `if price > discount ...` di repo | Pindahkan ke UseCase |
| Handler berisi business logic | Validasi kompleks di handler | Pindahkan ke UseCase |

---

*Terakhir diperbarui: 2026-05-03*
