# Go Starter Kit — Web API Template (Chi + Zerolog)

Starter kit ini adalah kerangka awal untuk membangun aplikasi backend berbasis Go (Web API) dengan struktur direktori rapi, logging modern (Zerolog), routing fleksibel (Chi), middleware standar, dan dukungan Docker. Tujuannya agar developer tidak perlu menyusun ulang struktur dari nol setiap kali memulai proyek baru.

---

## ✨ Fitur Utama

- **Struktur direktori standar** dengan pemisahan concern (cmd, internal, pkg)
- **Routing menggunakan Chi** + grouping per versi API
- **Middleware bawaan**: request logging, panic recovery, CORS
- **Logging** dengan Zerolog
- **Config via .env** (joho/godotenv)
- **Helper response JSON** standar
- **Contoh handler**: `/health`
- **Dockerfile multi-stage** & `docker-compose.yml`
- **Makefile** untuk run, build, test, lint, docker
- **Init script** (`init.sh`) untuk mengganti `MODULE_PATH`
- **Optional live reload** via [Air](https://github.com/cosmtrek/air)

---

## 📂 Struktur Direktori

```
├─ cmd/
│  └─ api/
│     └─ main.go          # Entrypoint aplikasi
├─ internal/
│  ├─ config/             # Load env & konfigurasi
│  ├─ http/               # Router & middleware
│  ├─ handler/            # HTTP handler (contoh: health)
│  ├─ service/            # Business logic
│  ├─ repository/         # Layer repository (db, memory, dll)
│  └─ logger/             # Logging setup
├─ pkg/
│  └─ response/           # Helper response JSON
├─ .env.example           # Contoh environment variable
├─ Makefile               # Task build/test/lint/run
├─ Dockerfile             # Docker build multi-stage
├─ docker-compose.yml     # Docker-compose (opsional + DB)
├─ go.mod                 # Go module
├─ init.sh                # Script inisialisasi modul
└─ .gitignore             # Ignore file
```

---

## 🚀 Quick Start

### 1. Clone / Gunakan Template

Clone atau gunakan repo ini sebagai template di GitHub.

```bash
git clone https://github.com/username/go-starter-kit myapp
cd myapp
```

### 2. Inisialisasi Modul

Jalankan script init untuk mengganti `MODULE_PATH` dengan modul project Anda:

```bash
./init.sh github.com/username/myapp
cp .env.example .env
```

### 3. Jalankan Aplikasi

```bash
make run
# atau
go run ./cmd/api
```

### 4. Cek Endpoint

```bash
curl http://localhost:8080/health
# → { "status": "ok", "message": "service healthy" }
```

---

## ⚙️ Konfigurasi

Semua konfigurasi dilakukan melalui environment variables. File `.env.example` bisa dijadikan acuan.

Contoh `.env`:

```env
APP_NAME=go-starter
APP_ENV=development
HTTP_PORT=8080
LOG_LEVEL=info
DATABASE_URL=postgres://app:app@localhost:5432/app?sslmode=disable
```

---

## 🛠 Makefile Commands

| Command            | Deskripsi                               |
| ------------------ | --------------------------------------- |
| `make run`         | Menjalankan server (`go run ./cmd/api`) |
| `make build`       | Build binary ke `bin/go-starter`        |
| `make test`        | Menjalankan unit test                   |
| `make tidy`        | Menjalankan `go mod tidy`               |
| `make lint`        | Menjalankan golangci-lint               |
| `make docker-up`   | Build & run docker-compose              |
| `make docker-down` | Hentikan docker-compose                 |

---

## 🐳 Docker

### Build & Run dengan Dockerfile

```bash
docker build -t go-starter .
docker run -p 8080:8080 --env-file .env go-starter
```

### Jalankan dengan Docker Compose

```bash
make docker-up
# API akan jalan di http://localhost:8080
```

---

## 📦 Dependency Utama

- [chi](https://github.com/go-chi/chi) — lightweight router
- [zerolog](https://github.com/rs/zerolog) — structured logging
- [godotenv](https://github.com/joho/godotenv) — env loader
- [uuid](https://github.com/google/uuid) — UUID generator

---

## 📐 Arsitektur

- **Handler (internal/handler):** menerima request HTTP, parsing input, kirim ke service.
- **Service (internal/service):** berisi business logic.
- **Repository (internal/repository):** abstraksi database atau storage.
- **Config (internal/config):** memuat konfigurasi aplikasi.
- **Logger (internal/logger):** setup zerolog sesuai level env.
- **Middleware (internal/http/middleware):** logging, recovery, CORS.

---

## 🔧 Extend Project

- Tambahkan **fitur baru** dengan membuat handler, service, repository.
- Gunakan **pgx/gorm** jika butuh database PostgreSQL.
- Tambahkan **validator** (go-playground/validator) untuk input request.
- Tambahkan **JWT middleware** untuk authentication.
- Versi API dipisahkan lewat route `/v1`, `/v2`, dst.

---

## ✅ Checklist Sebelum Deploy

- [ ] Semua unit test lulus: `make test`
- [ ] Linter & vet bersih: `make lint`
- [ ] Build docker sukses
- [ ] Endpoint `/health` mengembalikan 200 OK
- [ ] ENV sesuai dengan environment (dev/staging/prod)

---

## 📄 License

MIT — bebas digunakan, modifikasi, distribusi.

---

Selamat membangun dengan Go Starter Kit 🚀
