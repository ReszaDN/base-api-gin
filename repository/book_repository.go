package repository

import (
	"base-be/entity"
	"gorm.io/gorm"
)

// --- KONTRAK (INTERFACE) ---

// BookRepository adalah interface (kontrak) untuk repository buku.
// Service akan bergantung pada interface ini, bukan pada implementasi struct-nya.
type BookRepository interface {
	Create(book entity.Book) (entity.Book, error)
	FindAll() ([]entity.Book, error)
	FindByID(ID int) (entity.Book, error)
	Update(book entity.Book) (entity.Book, error)
	Delete(book entity.Book) error
}

// --- IMPLEMENTASI ---

// bookRepository adalah implementasi (struct) dari BookRepository
type bookRepository struct {
	db *gorm.DB // Dependency ke GORM
}

// NewBookRepository adalah constructor untuk bookRepository.
// Ini akan dipanggil di cmd/api/main.go (saat wiring)
func NewBookRepository(db *gorm.DB) *bookRepository {
	return &bookRepository{db: db}
}

// Implementasi method Create
func (r *bookRepository) Create(book entity.Book) (entity.Book, error) {
	err := r.db.Create(&book).Error
	return book, err
}

// Implementasi method FindAll
func (r *bookRepository) FindAll() ([]entity.Book, error) {
	var books []entity.Book
	err := r.db.Find(&books).Error
	return books, err
}

// Implementasi method FindByID
func (r *bookRepository) FindByID(ID int) (entity.Book, error) {
	var book entity.Book
	// Pakai 'First' agar GORM mengembalikan error 'record not found' jika kosong
	err := r.db.First(&book, ID).Error
	return book, err
}

// Implementasi method Update
func (r *bookRepository) Update(book entity.Book) (entity.Book, error) {
	err := r.db.Save(&book).Error // Save akan update semua field
	return book, err
}

// Implementasi method Delete
func (r *bookRepository) Delete(book entity.Book) error {
	err := r.db.Delete(&book).Error
	return err
}
