# 📄 Feature Document Template

> **Gunakan template ini untuk mendokumentasikan setiap fitur baru sebelum implementasi.**

---

## 🏷️ Feature: [Nama Fitur]

**Author**: [Nama]  
**Created**: [YYYY-MM-DD]  
**Status**: 📋 Planned | 🔨 In Progress | ✅ Done  
**Priority**: 🔴 High | 🟡 Medium | 🟢 Low  

---

## 📋 Deskripsi

[Jelaskan fitur ini secara singkat — apa yang dilakukan dan mengapa dibutuhkan]

---

## 🎯 User Stories

1. Sebagai [role], saya ingin [aksi], agar [manfaat]
2. ...

---

## 📐 Technical Design

### Domain Entity

```go
type [EntityName] struct {
    // fields
}
```

### Repository Interface

```go
type [Entity]Repository interface {
    // methods
}
```

### Use Case Interface

```go
type [Entity]UseCase interface {
    // methods
}
```

---

## 🔗 Endpoints

| Method | Endpoint | Deskripsi | Auth |
|--------|----------|-----------|------|
| GET | `/api/v1/...` | ... | ☐/☑ |
| POST | `/api/v1/...` | ... | ☐/☑ |

---

## 📊 Request/Response Examples

### Request

```json
{
    "field": "value"
}
```

### Response

```json
{
    "success": true,
    "data": {}
}
```

---

## ⚙️ Business Rules

1. [Rule 1]
2. [Rule 2]

---

## 🧪 Test Scenarios

| # | Scenario | Input | Expected Output |
|---|----------|-------|-----------------|
| 1 | Happy path | ... | ... |
| 2 | Validation error | ... | ... |
| 3 | Not found | ... | ... |

---

## 📁 Files to Create/Modify

- [ ] `internal/domain/[entity].go`
- [ ] `internal/usecase/[entity]_usecase.go`
- [ ] `internal/repository/mongo/[entity]_repository.go`
- [ ] `internal/delivery/http/handler/[entity]_handler.go`
- [ ] Test files for each layer

---

## 🔗 Dependencies

- Depends on: [list features this depends on]
- Depended by: [list features that depend on this]

---

## 📝 Notes

[Catatan tambahan, edge cases, atau keputusan arsitektur yang perlu diingat]
