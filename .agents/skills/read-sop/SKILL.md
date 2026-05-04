---
name: Read Project SOP
description: >
  Membaca semua Standard Operating Procedures (SOP) dan dokumentasi proyek sebelum
  merencanakan atau mengimplementasikan fitur apapun. Skill ini WAJIB dipanggil
  di awal setiap sesi planning atau coding fitur baru.
---

# Read Project SOP

## Kapan Skill Ini Digunakan

Skill ini **WAJIB** dijalankan ketika:
- Merencanakan fitur baru
- Memulai implementasi fitur
- Melakukan code review
- Membuat design document
- Menambah endpoint API baru
- Refactoring kode yang sudah ada

## Instruksi

1. **Baca file `RULES.md`** di root proyek — Ini adalah panduan utama yang berisi aturan global, tech stack, dan larangan-larangan penting.

2. **Baca SOP yang relevan** di `docs/SOP/`:
   - `01-development-workflow.md` — Alur kerja pengembangan
   - `02-code-standards.md` — Naming convention & code style
   - `03-git-branching-strategy.md` — Strategi Git
   - `04-code-review.md` — Checklist code review
   - `05-testing-strategy.md` — Strategi testing per layer
   - `06-api-design.md` — Standar desain API & Swagger
   - `07-deployment.md` — Deployment & CI/CD
   - `08-error-handling.md` — Error handling & logging

3. **Baca arsitektur** di `docs/architecture/`:
   - `clean-architecture.md` — Layer rules & dependency direction
   - `project-structure.md` — Folder structure & aturan per folder
   - `dependency-graph.md` — Dependency graph antar domain

4. **Cek status fitur** di `docs/features/feature-summary.md` — Pastikan fitur belum ada atau perlu diupdate.

5. **Cek endpoint** di `docs/api/endpoints.md` — Pastikan endpoint yang akan dibuat belum ada.

6. **Cek progress** di `docs/todo/master-todo.md` — Lihat task mana yang sedang/sudah dikerjakan.

## Output yang Diharapkan

Setelah membaca SOP, kamu harus bisa menjawab:
- Apa arsitektur yang digunakan? (Clean Architecture 4-layer)
- Apa tech stack-nya? (Go, MongoDB Atlas, Swagger/Swaggo)
- Bagaimana urutan implementasi fitur? (Domain → UseCase → Repository → Handler)
- Apa naming convention yang berlaku?
- Apa standar response format API?
- Fitur apa yang sudah ada dan apa yang belum?

## Peringatan

- ❌ JANGAN skip membaca SOP
- ❌ JANGAN gunakan Prisma Go (sudah deprecated)
- ❌ JANGAN buat kode yang melanggar clean architecture layer rules
- ❌ JANGAN buat endpoint tanpa Swagger annotation
