# Proyek Backend Go (base-be)

Ini adalah starter template (kerangka awal) untuk proyek backend Go, yang dibangun dengan **Gin**, **GORM**, dan arsitektur **Repository Pattern** yang mumpuni (scalable).

Struktur proyek ini dirancang untuk pemisahan tanggung jawab (*Separation of Concerns*) yang jelas, modularitas, dan kemudahan dalam pengujian (*testing*).

---

## ✨ Fitur Utama

- **Framework:** Gin (Router HTTP)
- **ORM:** GORM (ORM untuk PostgreSQL)
- **Konfigurasi:** Viper (Manajemen konfigurasi via file config.yaml)
- **Arsitektur:** Repository Pattern & Service Layer
- **Struktur Modular:** Rute dan logika bisnis dipisahkan berdasarkan domain/modul (misal: Buku, User).
- **Pemisahan Perintah:** Perintah terpisah untuk menjalankan server APInya (`cmd/api`) dan migrasi database (`cmd/migrate`).

---

## 🏛️ Arsitektur & Struktur Folder

Struktur ini dirancang untuk memisahkan logika berdasarkan tanggung jawabnya.

```
base-be/
│
├── cmd/                # (Entrypoints/Perintah Eksekusi)
│   ├── api/            # (Entrypoint untuk Server API Utama)
│   │   └── main.go     # (Tugas: Wiring DB, Repo, Service, Handler, Routes & Run Server)
│   └── migrate/        # (Entrypoint untuk Alat Migrasi DB)
│       └── main.go     # (Tugas: Load Config, Connect DB, Run GORM AutoMigrate)
│
├── config/             # (Manajemen Konfigurasi)
│   └── config.go       # (Logika Viper untuk memuat config.yaml/env)
│
├── dto/                # (Data Transfer Objects)
│   └── book_dto.go     # (Structs: BookRequest, BookResponse. Validasi 'binding')
│
├── entity/             # (Model/Entitas Database)
│   └── book.go         # (Struct 'Book' dengan GORM tags)
│
├── handler/            # (Lapisan HTTP)
│   └── book_handler.go # (Gin handlers untuk 'Book': PostBook, GetBook, etc.)
│
├── repository/         # (Lapisan Akses Data)
│   └── book_repository.go # (Implementasi query GORM untuk 'Book')
│
├── routes/             # (Definisi Rute/Endpoint)
│   ├── routes.go       # (File Hub: SetupRoutes, membuat grup /v1, memanggil rute modular)
│   └── book_routes.go  # (Modular: SetupBookRoutes, mendaftarkan /v1/books/...)
│
├── service/            # (Lapisan Logika Bisnis)
│   └── book_service.go # (Interface & Implementasi Logika Bisnis 'Book')
│
├── .gitignore
├── config.yaml.example # (Template/Contoh file konfigurasi - AMAN DI-COMMIT)
└── README.md
```

---

## 📂 Penjelasan Folder

- **cmd/** : Folder "Commands". Setiap subfolder adalah entrypoint yang bisa dieksekusi (`main.go`).
  - `api/` : Menjalankan server Gin.
  - `migrate/` : Menjalankan migrasi GORM (AutoMigrate).
- **config/** : Mengelola loading konfigurasi menggunakan Viper dari `config.yaml`.
- **entity/** : Mendefinisikan struct GORM (model database). Ini adalah "cetakan" untuk tabel database Anda.
- **dto/** : (*Data Transfer Object*) Mendefinisikan struct untuk data request (JSON input) dan response (JSON output).
- **handler/** : Lapisan HTTP. Bertugas menerima request dari Gin, mem-binding/validasi DTO, memanggil service, dan mengembalikan response JSON.
- **service/** : Lapisan Logika Bisnis. Otak dari aplikasi. Bertugas mengorkestrasi alur kerja, melakukan validasi bisnis, dan memanggil repository.
- **repository/** : Lapisan Akses Data. Bertugas untuk berbicara dengan database (GORM).
- **routes/** : Mendefinisikan endpoint API.

---

## 🚀 Panduan Memulai

Berikut adalah langkah-langkah untuk menjalankan proyek ini secara lokal.

### 1️⃣ Prasyarat

- **Go** (versi 1.18 atau lebih baru)
- **PostgreSQL** (Database)
- Sebuah database manager (opsional, seperti DBeaver atau Postico)

### 2️⃣ Instalasi

[//]: # (Clone repositori ini:)

[//]: # (```bash)

[//]: # (git clone [URL_GIT_ANDA_DI_SINI] base-be)

[//]: # (cd base-be)

[//]: # (```)

Install Go :

```bash
go mod init [NAMA_MODUL] ( ex. base-be )
```

Install dependensi:
```bash
go mod tidy
```

### 3️⃣ Konfigurasi

Proyek ini menggunakan file `config.yaml` untuk manajemen konfigurasi (seperti koneksi DB).

Buat file config.yaml:
```bash
cp config.yaml.example config.yaml
```

[//]: # (*&#40;Catatan: Anda perlu membuat file `config.yaml.example` terlebih dahulu sebagai template&#41;*)

Edit config.yaml:
```yaml
server:
  port: "8080"

database:
  host: "localhost"
  port: "5432"
  user: "postgres"
  password: "password_anda"
  dbname: "nama_database_anda"
```

---

### 4️⃣ Menjalankan Migrasi Database

Proyek ini memiliki perintah terpisah untuk migrasi.
Jalankan perintah ini setiap kali Anda mengubah struktur entity:

```bash
go run ./cmd/migrate
```

Output yang diharapkan:
```
Menjalankan migrasi...
Migrasi berhasil diselesaikan.
```

---

### 5️⃣ Menjalankan Server API

Setelah database dan tabel siap, jalankan server API utama:

```bash
go run ./cmd/api
```

Output yang diharapkan:
```
Menjalankan server API di port 8080
```

Server Anda sekarang berjalan di:
👉 http://localhost:8080

---

## 📦 Contoh API Endpoints

Semua endpoint diawali dengan `/v1`.

| Method | Endpoint     | Deskripsi                    |
|---------|--------------|------------------------------|
| POST    | /books       | Membuat buku baru            |
| GET     | /books       | Mendapatkan semua buku       |
| GET     | /books/:id   | Mendapatkan buku berdasarkan ID |
| PUT     | /books/:id   | Memperbarui buku berdasarkan ID |
| DELETE  | /books/:id   | Menghapus buku berdasarkan ID |

---
