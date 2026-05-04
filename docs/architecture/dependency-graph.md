# Dependency Graph & AI Context

> **Tujuan**: Memberikan konteks arsitektur kepada AI assistant dan developer melalui dependency graph.

---

## 🔗 Dependency Graph

```mermaid
graph TD
    subgraph "Entry Point"
        MAIN[cmd/api/main.go]
    end
    
    subgraph "Delivery Layer"
        ROUTER[router.go]
        AUTH_H[auth_handler.go]
        USER_H[user_handler.go]
        PROD_H[product_handler.go]
        CART_H[cart_handler.go]
        ORDER_H[order_handler.go]
        AUTH_MW[auth_middleware.go]
        CORS_MW[cors_middleware.go]
        LOG_MW[logging_middleware.go]
    end
    
    subgraph "UseCase Layer"
        AUTH_UC[auth_usecase.go]
        USER_UC[user_usecase.go]
        PROD_UC[product_usecase.go]
        CART_UC[cart_usecase.go]
        ORDER_UC[order_usecase.go]
    end
    
    subgraph "Repository Layer"
        USER_R[user_repository.go]
        PROD_R[product_repository.go]
        CART_R[cart_repository.go]
        ORDER_R[order_repository.go]
    end
    
    subgraph "Domain Layer - Core"
        USER_E[user.go]
        PROD_E[product.go]
        CART_E[cart.go]
        ORDER_E[order.go]
        ERRORS[errors.go]
        DTO[dto.go]
    end
    
    subgraph "Infrastructure"
        MONGO[mongodb/connection.go]
        JWT[token/jwt.go]
        HASH[hash/bcrypt.go]
        CONFIG[config/config.go]
    end
    
    MAIN --> ROUTER
    MAIN --> CONFIG
    MAIN --> MONGO
    
    ROUTER --> AUTH_H & USER_H & PROD_H & CART_H & ORDER_H
    ROUTER --> AUTH_MW & CORS_MW & LOG_MW
    
    AUTH_H --> AUTH_UC
    USER_H --> USER_UC
    PROD_H --> PROD_UC
    CART_H --> CART_UC
    ORDER_H --> ORDER_UC
    
    AUTH_UC --> USER_E & ERRORS & JWT & HASH
    USER_UC --> USER_E & ERRORS
    PROD_UC --> PROD_E & ERRORS
    CART_UC --> CART_E & PROD_E & ERRORS
    ORDER_UC --> ORDER_E & CART_E & ERRORS
    
    USER_R --> USER_E & MONGO
    PROD_R --> PROD_E & MONGO
    CART_R --> CART_E & MONGO
    ORDER_R --> ORDER_E & MONGO
```

---

## 🤖 AI Context Guide

### Ketika AI merencanakan fitur baru, AI HARUS:

1. **Baca SOP terlebih dahulu** — `docs/SOP/` 
2. **Check Feature Summary** — `docs/features/feature-summary.md`
3. **Pahami Dependency Graph** — Diagram di atas
4. **Ikuti Clean Architecture** — `docs/architecture/clean-architecture.md`
5. **Gunakan Naming Convention** — `docs/SOP/02-code-standards.md`

### Context Map per Domain

| Domain | Depends On | Depended By |
|--------|-----------|-------------|
| **User** | - | Auth, Order, Review |
| **Product** | Category | Cart, Order, Review |
| **Category** | - | Product |
| **Cart** | Product, User | Order |
| **Order** | Cart, Product, User, Payment | - |
| **Payment** | Order | - |
| **Review** | Product, User, Order | Product (rating) |

### Urutan Implementasi Fitur (Recommended)

```mermaid
flowchart LR
    A[1. User & Auth] --> B[2. Category]
    B --> C[3. Product]
    C --> D[4. Cart]
    D --> E[5. Order]
    E --> F[6. Payment]
    C --> G[7. Review]
```

---

## 🔧 Tools untuk Generate Graph

```bash
# Install godepgraph
go install github.com/kisielk/godepgraph@latest

# Generate dependency graph
godepgraph -s ./cmd/api | dot -Tpng -o docs/dep-graph.png

# Install go-callvis untuk call graph
go install github.com/ofabry/go-callvis@latest
go-callvis -group pkg ./cmd/api
```

---

*Terakhir diperbarui: 2026-05-03*
