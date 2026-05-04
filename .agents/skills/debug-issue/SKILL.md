---
name: Debug Issue
description: >
  Skill untuk debugging dan troubleshooting issues di proyek E-Commerce Backend.
  Mengikuti pendekatan sistematis dari identifikasi, analisis, fix, hingga
  regression test.
---

# Debug Issue

## Kapan Skill Ini Digunakan

- Saat ada bug report
- Saat test gagal
- Saat ada error di runtime
- Saat ada performance issue

## Alur Debugging

### Step 1: Reproduce & Identify

1. Baca error message / bug report dengan teliti
2. Identifikasi layer mana yang bermasalah:
   - **Handler** — Masalah parsing, routing, response format
   - **UseCase** — Masalah business logic, orchestration
   - **Repository** — Masalah query, connection, data format
   - **Domain** — Masalah validation, entity rules
3. Cek file yang relevan berdasarkan dependency graph di `docs/architecture/dependency-graph.md`

### Step 2: Analyze

1. Baca kode di layer yang bermasalah
2. Trace data flow dari handler → usecase → repository → database
3. Cek error handling — apakah error di-wrap dengan benar?
4. Cek context propagation — apakah context timeout sudah benar?
5. Cek database query — apakah filter/index sudah benar?

### Step 3: Fix

1. **Tulis test yang mereproduksi bug** (regression test)
2. Perbaiki kode
3. Pastikan test baru passed
4. Pastikan test lama tetap passed
5. Follow code standards (SOP 02)

### Step 4: Verify

```bash
# Run all tests
go test ./internal/... -v -race -count=1

# Run specific test
go test ./internal/usecase/... -run TestSpecific -v

# Lint check
golangci-lint run
```

### Step 5: Document

- Update feature doc jika ada perubahan behavior
- Commit dengan format: `fix(scope): description of fix`

## Common Issues & Solutions

| Issue | Layer | Kemungkinan Penyebab |
|-------|-------|---------------------|
| 404 Not Found | Handler/Router | Route belum di-register |
| 400 Bad Request | Handler | Request body parsing gagal |
| 401 Unauthorized | Middleware | Token invalid/expired |
| 500 Internal Error | Any | Unhandled error |
| Empty response | Repository | Query filter salah |
| Slow query | Repository | Missing index |
| Data inconsistency | UseCase | Business logic error |

## Rules

- ✅ Selalu tulis regression test SEBELUM fix
- ✅ Fix di layer yang tepat (jangan patch di handler jika masalah di usecase)
- ✅ Gunakan `fmt.Errorf("context: %w", err)` untuk tracing
- ❌ Jangan asal fix tanpa memahami root cause
- ❌ Jangan skip testing setelah fix
