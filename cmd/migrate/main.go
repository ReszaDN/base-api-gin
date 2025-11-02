package main

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	// Import yang kita perlukan
	"base-be/config"
	"base-be/entity" // <-- Kita perlu ini untuk struct model
)

func main() {
	log.Println("Menjalankan migrasi...")

	// 1. Load Config (kode duplikat, tapi tidak masalah di sini)
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Tidak bisa memuat file konfigurasi: %v", err)
	}

	// 2. Koneksi DB
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		cfg.Database.Host, cfg.Database.User, cfg.Database.Password, cfg.Database.DBName, cfg.Database.Port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB Connection Error")
	}

	// 3. INI DIA INTI-nya
	// Daftarkan semua model Anda di sini
	err = db.AutoMigrate(
		&entity.Book{},
		// &entity.User{},  // <-- Tambahkan model lain jika ada
		// &entity.Order{}, // <-- Seperti ini
	)

	if err != nil {
		log.Fatalf("Gagal melakukan migrasi: %v", err)
	}

	log.Println("Migrasi berhasil diselesaikan.")
}

// COMMAND UNTUK RUN MIGRATE : go run ./cmd/migrate
