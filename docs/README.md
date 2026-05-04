# 📚 E-Commerce Backend — Documentation Hub

> **Pusat dokumentasi untuk proyek E-Commerce Backend.**  
> Semua SOP, arsitektur, dan panduan pengembangan ada di sini.

---

## 📁 Struktur Dokumentasi

```
docs/
├── README.md                          # File ini — Daftar isi dokumentasi
├── SOP/
│   ├── 01-development-workflow.md     # SOP Alur Kerja Pengembangan
│   ├── 02-code-standards.md           # SOP Standar Kode & Naming Convention
│   ├── 03-git-branching-strategy.md   # SOP Strategi Git & Branching
│   ├── 04-code-review.md             # SOP Code Review
│   ├── 05-testing-strategy.md         # SOP Strategi Testing
│   ├── 06-api-design.md              # SOP Desain API & Endpoint
│   ├── 07-deployment.md              # SOP Deployment & CI/CD
│   └── 08-error-handling.md          # SOP Error Handling & Logging
├── architecture/
│   ├── clean-architecture.md          # Arsitektur Clean Architecture
│   ├── project-structure.md           # Struktur Proyek & Folder
│   └── dependency-graph.md            # Dependency Graph & AI Context
├── features/
│   ├── feature-summary.md             # Ringkasan Semua Fitur
│   └── feature-template.md            # Template Dokumen Fitur Baru
├── api/
│   ├── endpoints.md                   # Daftar Lengkap Endpoint API
│   └── swagger-guide.md              # Panduan Setup & Penggunaan Swagger
└── todo/
    └── master-todo.md                 # Master To-Do List Proyek
```

---

## 🚀 Tech Stack

| Komponen       | Teknologi                                    |
| -------------- | -------------------------------------------- |
| **Bahasa**     | Go (Golang) — Pure, tanpa framework berat    |
| **Database**   | MongoDB Atlas (Cloud)                        |
| **ORM/Driver** | Official MongoDB Go Driver (`mongo-go-driver`) |
| **API Docs**   | Swagger (Swaggo)                             |
| **AI Context** | Code Review Graph (MCP-based)                |
| **Arsitektur** | Clean Architecture (4-Layer)                 |

> ⚠️ **Catatan Penting tentang Prisma**: Prisma Client Go telah **di-archive** pada 2025 dan tidak lagi dimaintain. Sebagai gantinya, proyek ini menggunakan **Official MongoDB Go Driver** (`go.mongodb.org/mongo-driver`) yang lebih stabil, performa lebih baik, dan didukung penuh oleh MongoDB Inc.

---

## 🔗 Quick Links

| Dokumen | Deskripsi |
| ------- | --------- |
| [SOP Development Workflow](./SOP/01-development-workflow.md) | Alur kerja dari ideasi hingga deployment |
| [Clean Architecture](./architecture/clean-architecture.md) | Penjelasan arsitektur & layer dependencies |
| [Project Structure](./architecture/project-structure.md) | Penjelasan setiap folder & file |
| [Feature Summary](./features/feature-summary.md) | Ringkasan fitur E-Commerce |
| [API Endpoints](./api/endpoints.md) | Daftar lengkap REST API |
| [Master To-Do](./todo/master-todo.md) | Checklist pengembangan |

---

## 📖 Cara Menggunakan Dokumentasi Ini

1. **Sebelum mulai koding**: Baca [SOP Development Workflow](./SOP/01-development-workflow.md)
2. **Saat membuat fitur baru**: Gunakan [Feature Template](./features/feature-template.md) 
3. **Saat code review**: Ikuti [SOP Code Review](./SOP/04-code-review.md)
4. **Saat desain API**: Ikuti [SOP API Design](./SOP/06-api-design.md)
5. **AI Assistant**: Selalu baca folder `docs/` sebelum merencanakan fitur

---

*Terakhir diperbarui: 2026-05-03*
