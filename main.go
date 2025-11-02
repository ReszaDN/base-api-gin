package main

import (
	"fmt" // 1. Tambahkan import "fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"base-be/config" // 2. Import paket config baru kita
)

func main() {

	// 3. Panggil LoadConfig() di paling atas
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Tidak bisa memuat file konfigurasi: %v", err)
	}

	// 4. HAPUS DSN YANG LAMA
	// dsn := "host=localhost user=postgres password=redn8222 dbname=coba-golang port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	// 5. BUAT DSN BARU dari struct 'cfg' yang sudah di-load
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		cfg.Database.Host,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.DBName,
		cfg.Database.Port,
	)

	// 6. Koneksikan GORM (kode Anda selanjutnya tidak berubah)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB Connection Error")
	}

	db.AutoMigrate(&book.Book{})
	
	r := gin.Default()

	// 7. (BONUS) Jalankan server menggunakan port dari config
	log.Printf("Menjalankan server di port %s", cfg.Server.Port)
	r.Run(":" + cfg.Server.Port) // Ganti dari r.Run() saja
}
