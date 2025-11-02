package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	// Perhatikan path import sekarang berubah
	"base-be/config"
	"base-be/handler"
	"base-be/repository"
	"base-be/routes"
	"base-be/service"
	// "base-be/entity" // <-- Hapus import entity/book
)

func main() {
	// 1. Load Config
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

	// 3. Dependency Injection (Wiring)
	bookRepository := repository.NewBookRepository(db)
	bookService := service.NewBookService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)
	// userHandler := handler.NewUserHandler(userService) // (Ini untuk nanti)

	// 4. Init Router & Routes
	r := gin.Default()

	// 5. Run Server
	// main.go tidak perlu tahu seberapa kompleks struktur rute Anda.
	routes.SetupRoutes(r, bookHandler) // (Nanti: , userHandler)

	log.Printf("Menjalankan server API di port %s", cfg.Server.Port)
	r.Run(":" + cfg.Server.Port)
}

// COMMAND UNTUK RUN PROJECT : go run ./cmd/api
