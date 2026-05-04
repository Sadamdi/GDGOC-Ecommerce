# 📋 Master To-Do List — E-Commerce Backend

> **Checklist pengembangan proyek E-Commerce Backend.**  
> Update status setiap kali ada progress.

---

## Phase 0: Project Setup ⚙️

- [ ] Initialize Go module (`go mod init`)
- [ ] Setup project folder structure (clean architecture)
- [ ] Setup MongoDB Atlas connection
- [ ] Create configuration management (env vars)
- [ ] Setup Makefile
- [ ] Setup `.gitignore` & `.env.example`
- [ ] Setup `golangci-lint` configuration
- [ ] Setup Swagger (Swaggo)
- [ ] Setup Dockerfile & docker-compose
- [ ] Setup basic middleware (CORS, logging, recovery)
- [ ] Setup standard response helpers

---

## Phase 1: Authentication & User Management 🔐

- [ ] **Domain Layer**
  - [ ] User entity
  - [ ] UserRepository interface
  - [ ] AuthUseCase interface
  - [ ] Domain errors
  - [ ] Request/Response DTOs

- [ ] **Repository Layer**
  - [ ] MongoDB User repository implementation
  - [ ] Create indexes (email unique)

- [ ] **UseCase Layer**
  - [ ] Register use case
  - [ ] Login use case
  - [ ] Refresh token use case
  - [ ] Password hashing (bcrypt)
  - [ ] JWT token generation

- [ ] **Handler Layer**
  - [ ] Register handler + Swagger
  - [ ] Login handler + Swagger
  - [ ] Refresh handler + Swagger
  - [ ] Auth middleware (JWT validation)

- [ ] **Testing**
  - [ ] Domain validation tests
  - [ ] UseCase unit tests (mock repo)
  - [ ] Handler HTTP tests
  - [ ] Integration tests

---

## Phase 2: Product Catalog 📦

- [ ] **Domain Layer**
  - [ ] Product entity
  - [ ] Category entity
  - [ ] ProductRepository interface
  - [ ] CategoryRepository interface

- [ ] **Repository Layer**
  - [ ] MongoDB Product repository
  - [ ] MongoDB Category repository
  - [ ] Search & filter implementation
  - [ ] Pagination implementation

- [ ] **UseCase Layer**
  - [ ] Product CRUD use cases
  - [ ] Category CRUD use cases
  - [ ] Product search & filter

- [ ] **Handler Layer**
  - [ ] Product handlers + Swagger
  - [ ] Category handlers + Swagger

- [ ] **Testing**
  - [ ] Unit tests
  - [ ] Integration tests

---

## Phase 3: Shopping Cart 🛒

- [ ] **Domain Layer**
  - [ ] Cart entity
  - [ ] CartRepository interface

- [ ] **Repository Layer**
  - [ ] MongoDB Cart repository

- [ ] **UseCase Layer**
  - [ ] Add/Remove/Update cart items
  - [ ] Stock validation
  - [ ] Cart total calculation

- [ ] **Handler Layer**
  - [ ] Cart handlers + Swagger

- [ ] **Testing**
  - [ ] Unit tests
  - [ ] Integration tests

---

## Phase 4: Order Management 📑

- [ ] **Domain Layer**
  - [ ] Order entity
  - [ ] OrderRepository interface
  - [ ] Order status enum & state machine

- [ ] **Repository Layer**
  - [ ] MongoDB Order repository

- [ ] **UseCase Layer**
  - [ ] Checkout (cart → order)
  - [ ] Order status management
  - [ ] Stock deduction on order

- [ ] **Handler Layer**
  - [ ] Order handlers + Swagger

- [ ] **Testing**
  - [ ] Unit tests
  - [ ] Integration tests

---

## Phase 5: Payment 💳

- [ ] **Domain Layer**
  - [ ] Payment entity
  - [ ] PaymentRepository interface

- [ ] **Repository & UseCase**
  - [ ] Payment processing
  - [ ] Payment status tracking

- [ ] **Handler Layer**
  - [ ] Payment handlers + Swagger

---

## Phase 6: Reviews & Admin ⭐🛡️

- [ ] **Reviews**
  - [ ] Review CRUD
  - [ ] Rating calculation

- [ ] **Admin Panel**
  - [ ] User management
  - [ ] Order management
  - [ ] Dashboard statistics

---

## Phase 7: Polish & Optimization 🔧

- [ ] Rate limiting middleware
- [ ] Request validation middleware
- [ ] API versioning
- [ ] Database indexing optimization
- [ ] Error response standardization
- [ ] Performance testing
- [ ] Security audit
- [ ] Documentation review
- [ ] CI/CD pipeline setup

---

## 📈 Progress Tracker

| Phase | Status | Progress |
|-------|--------|----------|
| Phase 0: Setup | 📋 Not Started | 0% |
| Phase 1: Auth & User | 📋 Not Started | 0% |
| Phase 2: Products | 📋 Not Started | 0% |
| Phase 3: Cart | 📋 Not Started | 0% |
| Phase 4: Orders | 📋 Not Started | 0% |
| Phase 5: Payment | 📋 Not Started | 0% |
| Phase 6: Reviews & Admin | 📋 Not Started | 0% |
| Phase 7: Polish | 📋 Not Started | 0% |

---

*Terakhir diperbarui: 2026-05-03*
