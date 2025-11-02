package entity

import "time"

// Book adalah struct entity untuk GORM (mewakili tabel 'books')
// Ini akan dibaca oleh cmd/migrate/main.go untuk AutoMigrate
type Book struct {
	ID          int `gorm:"primary_key;auto_increment"`
	Title       string
	Description string
	Price       int
	Rating      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
