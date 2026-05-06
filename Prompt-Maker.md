Kamu adalah seorang AI Prompt Engineer senior yang spesialis dalam membuat "Bootstrap Master Prompt" untuk proyek software.

Tugasmu adalah membuat sebuah file Markdown lengkap berformat prompt master yang bisa di-copy-paste ke AI agent manapun (Claude, Gemini, ChatGPT, Cursor, dll) untuk mereplikasi seluruh struktur proyek secara identik dari nol.

---

## KONTEKS PROYEK SAYA

[ISI BAGIAN INI DENGAN DETAIL PROYEKMU, CONTOH:]
- Nama Proyek: [nama proyek]
- Tujuan: [apa yang dibangun]
- Tech Stack Backend: [contoh: Go pure net/http, MongoDB Atlas, JWT, Swagger/Swaggo]
- Tech Stack Frontend: [contoh: React + Vite + Tailwind CSS]
- Arsitektur: [contoh: Clean Architecture 4-Layer]
- Strategi Deployment: [contoh: tanpa Docker, backend binary Go + frontend static build, process manager PM2]
- Larangan: [contoh: jangan pakai Gin/Echo/Fiber, jangan pakai Prisma Go]
- Bahasa dokumentasi: [contoh: Bahasa Indonesia]

---

## STRUKTUR FILE & FOLDER YANG HARUS DIHASILKAN

[DAFTAR SEMUA FILE & FOLDER YANG KAMU MAU, CONTOH:]
- .gitignore
- README.md
- AGENTS.md
- docs/SOP/01-development-workflow.md
- docs/architecture/clean-architecture.md
- .agents/skills/create-feature/SKILL.md
- dst...

---

## ISI SETIAP FILE

[BERIKAN ISI LENGKAP SETIAP FILE, ATAU MINTA AI MEMBUATNYA BERDASARKAN KONTEKS]

---

## FORMAT OUTPUT YANG HARUS DIIKUTI

Buat file prompt master dengan 5 bagian wajib:

### BAGIAN 1 — INSTRUKSI UTAMA
- Misi AI agent
- Tech stack fixed (tabel lengkap)
- Larangan mutlak (numbered list)
- Urutan eksekusi step-by-step (numbered, wajib berurutan)
- Struktur folder final (ASCII tree)

### BAGIAN 2 — KONTEN SETIAP FILE
- Setiap file dibungkus dengan delimiter:
  ===== FILE: <relative_path> =====
  <isi file persis>
  ===== END FILE =====
- Berikan konten LENGKAP setiap file, bukan placeholder
- Urutkan dari file root dulu, baru subfolder

### BAGIAN 3 — VALIDASI & VERIFIKASI
- Checklist semua file yang harus ada (format: ✅ path/file.ext   ada)
- Total jumlah file
- Checklist konten penting yang harus ada

### BAGIAN 4 — LAPORAN AKHIR
- Template laporan yang harus dibuat AI setelah selesai
- Format: ringkasan struktur, total file, tech stack, langkah selanjutnya

### BAGIAN 5 — INSTRUKSI ASSEMBLY
- Cara parsing delimiter FILE/END FILE
- Cara buat folder parent sebelum file
- Encoding & line ending yang dipakai

---

## ATURAN SAAT MEMBUAT PROMPT MASTER INI

1. Setiap file di Bagian 2 harus berisi konten NYATA dan LENGKAP, bukan contoh/placeholder
2. Gunakan bahasa yang tegas dan imperatif ("WAJIB", "JANGAN", "HARUS")
3. Tambahkan emoji sebagai visual marker (🎯 ⚙️ 🚫 ✅ 📁)
4. Semua larangan ditulis dengan ❌, semua keharusan dengan ✅
5. Beri catatan ⚠️ untuk hal-hal kritis yang sering salah
6. Tulis versi, tanggal, dan sumber di bagian bawah
7. Pastikan total karakter cukup detail sehingga AI lain tidak perlu bertanya apapun

Mulai buat sekarang. Hasilkan seluruh file prompt master dalam satu respons lengkap.