\# Proyek Backend Go (base-be)



Ini adalah \*starter template\* (kerangka awal) untuk proyek backend Go, yang dibangun dengan Gin, GORM, dan arsitektur Repository Pattern yang mumpuni (\*scalable\*).



Struktur proyek ini dirancang untuk pemisahan tanggung jawab (\*Separation of Concerns\*) yang jelas, modularitas, dan kemudahan dalam pengujian (\*testing\*).



---



\## ✨ Fitur Utama



\* \*\*Framework:\*\* \[Gin](https://github.com/gin-gonic/gin) (Router HTTP)

\* \*\*ORM:\*\* \[GORM](https://gorm.io/) (ORM untuk PostgreSQL)

\* \*\*Konfigurasi:\*\* \[Viper](https://github.com/spf13/viper) (Manajemen konfigurasi via file `config.yaml`)

\* \*\*Arsitektur:\*\* Repository Pattern \& Service Layer

\* \*\*Struktur Modular:\*\* Rute dan logika bisnis dipisahkan berdasarkan domain/modul (misal: Buku, User).

\* \*\*Pemisahan Perintah:\*\* Perintah terpisah untuk menjalankan server API (`cmd/api`) dan migrasi database (`cmd/migrate`).



---





\## Arsitektur \& Struktur Folder
Struktur ini dirancang untuk memisahkan logika berdasarkan tanggung jawabnya.



```bash

base-be/

│

├── cmd/                # (Entrypoints/Perintah Eksekusi)

│   ├── api/            # (Entrypoint untuk Server API Utama)

│   │   └── main.go     # (Tugas: Wiring DB, Repo, Service, Handler, Routes \& Run Server)

│   └── migrate/        # (Entrypoint untuk Alat Migrasi DB)

│       └── main.go     # (Tugas: Load Config, Connect DB, Run GORM AutoMigrate)

│

├── config/             # (Manajemen Konfigurasi)

│   └── config.go       # (Logika Viper untuk memuat config.yaml/env)

│

├── dto/                # (Data Transfer Objects)

│   └── book\_dto.go     # (Structs: BookRequest, BookResponse. Validasi 'binding')

│

├── entity/             # (Model/Entitas Database)

│   └── book.go         # (Struct 'Book' dengan GORM tags)

│

├── handler/            # (Lapisan HTTP)

│   └── book\_handler.go # (Gin handlers untuk 'Book': PostBook, GetBook, etc.)

│

├── repository/         # (Lapisan Akses Data)

│   └── book\_repository.go # (Implementasi query GORM untuk 'Book')

│

├── routes/             # (Definisi Rute/Endpoint)

│   ├── routes.go       # (File Hub: SetupRoutes, membuat grup /v1, memanggil rute modular)

│   └── book\_routes.go  # (Modular: SetupBookRoutes, mendaftarkan /v1/books/...)

│

├── service/            # (Lapisan Logika Bisnis)

│   └── book\_service.go # (Interface \& Implementasi Logika Bisnis 'Book')

│

├── .gitignore

├── config.yaml         # (File konfigurasi aktual - DIABAIKAN OLEH GIT)

├── config.yaml.example # (Template/Contoh file konfigurasi - AMAN DI-COMMIT)

├── go.mod

├── go.sum

└── README.md

```



\### Penjelasan Folder



\* \*\*`cmd/`\*\*: Folder "Commands". Setiap subfolder adalah \*entrypoint\* yang bisa dieksekusi (`main.go`).

&nbsp;   \* `api/`: Menjalankan server Gin.

&nbsp;   \* `migrate/`: Menjalankan migrasi GORM (`AutoMigrate`).

\* \*\*`config/`\*\*: Mengelola \*loading\* konfigurasi menggunakan Viper dari `config.yaml`.

\* \*\*`entity/`\*\*: Mendefinisikan \*struct\* GORM (model database). Ini adalah "cetakan" untuk tabel database Anda.

\* \*\*`dto/`\*\*: (Data Transfer Object) Mendefinisikan \*struct\* untuk data \*request\* (JSON input) dan \*response\* (JSON output). Di sinilah \*tag\* validasi `binding:"required"` diletakkan.

\* \*\*`handler/`\*\*: Lapisan HTTP. Bertugas menerima \*request\* dari Gin, mem-binding/validasi DTO, memanggil \*service\*, dan mengembalikan \*response\* JSON. \*\*Tidak boleh ada logika bisnis di sini.\*\*

\* \*\*`service/`\*\*: Lapisan Logika Bisnis. \*\*Otak\*\* dari aplikasi. Bertugas mengorkestrasi alur kerja, melakukan validasi bisnis, dan memanggil \*repository\*.

\* \*\*`repository/`\*\*: Lapisan Akses Data. Bertugas untuk berbicara dengan database (GORM). \*\*Hanya berisi \*query\*\*\* (Create, Read, Update, Delete).

\* \*\*`routes/`\*\*: Mendefinisikan \*endpoint\* API.

&nbsp;   \* `routes.go` adalah \*hub\* yang mendelegasikan ke `\_routes.go` modular.

&nbsp;   \* `book\_routes.go` mendaftarkan semua rute yang terkait dengan `/v1/books`.



\## 🚀 Panduan Memulai



Berikut adalah langkah-langkah untuk menjalankan proyek ini secara lokal.



\### 1. Prasyarat



\* \[Go](https://golang.org/dl/) (versi 1.18 atau lebih baru)

\* \[PostgreSQL](https://www.postgresql.org/download/) (Database)

\* Sebuah \*database manager\* (opsional, seperti DBeaver atau Postico)



\### 2. Instalasi



1\.  \*\*Clone\*\* repositori ini:

&nbsp;   ```bash

&nbsp;   git clone \[URL\_GIT\_ANDA\_DI\_SINI] base-be

&nbsp;   cd base-be

&nbsp;   ```



2\.  \*\*Install Dependensi\*\*

&nbsp;   Go akan mengunduh semua \*package\* yang dibutuhkan secara otomatis:

&nbsp;   ```bash

&nbsp;   go mod tidy

&nbsp;   ```



\### 3. Konfigurasi



Proyek ini menggunakan file `config.yaml` untuk manajemen konfigurasi (seperti koneksi DB).



1\.  \*\*Buat file `config.yaml`\*\*

&nbsp;   Salin dari file contoh (Jangan pernah \*commit\* `config.yaml` asli Anda ke Git):

&nbsp;   ```bash

&nbsp;   cp config.yaml.example config.yaml

&nbsp;   ```

&nbsp;   \*(Catatan: Anda perlu membuat file `config.yaml.example` terlebih dahulu sebagai template)\*



2\.  \*\*Edit `config.yaml`\*\*

&nbsp;   Buka file `config.yaml` dan sesuaikan dengan pengaturan database PostgreSQL lokal Anda:

&nbsp;   ```yaml

&nbsp;   server:

&nbsp;     port: "8080"



&nbsp;   database:

&nbsp;     host: "localhost"

&nbsp;     port: "5432"

&nbsp;     user: "postgres"

&nbsp;     password: "password\_anda"

&nbsp;     dbname: "nama\_database\_anda"

&nbsp;   ```



\### 4. Menjalankan Migrasi Database



Proyek ini memiliki perintah terpisah untuk migrasi. Ini akan membaca `entity` Anda dan membuat tabel di database menggunakan GORM `AutoMigrate`.



Jalankan perintah ini \*\*setiap kali Anda mengubah struktur `entity`\*\*:

```bash

go run ./cmd/migrate
```



Output yang diharapkan:
Menjalankan migrasi...

Migrasi berhasil diselesaikan.





\### 5. Menjalankan Server API



Setelah database dan tabel siap, Anda bisa menjalankan server API utama:

```bash
go run ./cmd/api

```

Output yang diharapkan:

```bash

Menjalankan server API di port 8080

```









