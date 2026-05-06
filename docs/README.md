# 📚 GDGOC E-Commerce — Documentation Hub

> **Pusat dokumentasi untuk proyek E-Commerce Full-Stack.**  
> Study Jam 2026 — GDGOC UIN Maulana Malik Ibrahim Malang

---

## 📁 Struktur Dokumentasi

```
docs/
├── README.md                          # File ini — Daftar isi dokumentasi
├── assets/                            # Logo & media assets
│   └── gdgoc-uin-malang-logo.png     # Logo GDGOC UIN Malang
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

### Backend (Go)

| Komponen | Teknologi |
|----------|-----------|
| **Language** | Go (Golang) — Pure `net/http`, tanpa framework berat |
| **Database** | MongoDB Atlas (Cloud) |
| **DB Driver** | Official MongoDB Go Driver (`go.mongodb.org/mongo-driver`) |
| **API Docs** | Swagger (Swaggo) |
| **Auth** | JWT + bcrypt |
| **Linting** | golangci-lint |
| **Formatting** | gofmt, goimports |

### Frontend (React)

| Komponen | Teknologi |
|----------|-----------|
| **Framework** | React (via Vite) |
| **Language** | JavaScript / TypeScript |
| **Styling** | CSS Modules / Tailwind CSS |
| **HTTP Client** | Axios / Fetch API |
| **Routing** | React Router |
| **State** | React Context / Zustand |
| **Linting** | ESLint |
| **Formatting** | Prettier |

### AI & DevOps

| Komponen | Teknologi |
|----------|-----------|
| **AI Code Review** | [Code Review Graph](https://github.com/tirth8205/code-review-graph) (MCP-based) |
| **AI Assistant** | Claude / Gemini via MCP |
| **Local Runtime** | Native Go & Node.js commands |

> ⚠️ **Catatan Penting tentang Prisma**: Prisma Client Go telah **di-archive** pada 2025 dan tidak lagi dimaintain. Proyek ini menggunakan **Official MongoDB Go Driver** (`go.mongodb.org/mongo-driver`) yang lebih stabil dan didukung penuh oleh MongoDB Inc.

---

## 🔗 Quick Links

| Dokumen | Deskripsi |
|---------|-----------|
| [SOP Development Workflow](./SOP/01-development-workflow.md) | Alur kerja dari ideasi hingga deployment |
| [Code Standards](./SOP/02-code-standards.md) | Naming convention & code style |
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

## 👥 Team

| Member | Role | GitHub |
|--------|------|--------|
| Sulthan Adam Rahmadi | Project Manager & Backend | [@Sadamdi](https://github.com/Sadamdi) |
| Villhze | Backend Developer | [@Villhze](https://github.com/Villhze) |
| Rahmat Rafii | Backend Developer | [@rahmatrafii](https://github.com/rahmatrafii) |
| M. Fahd Khulloh | UI/UX & Frontend | [@MFahdKhulloh-221](https://github.com/MFahdKhulloh-221) |

---

*Terakhir diperbarui: 2026-05-04*
