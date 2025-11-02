package service

import (
	"base-be/dto"
	"base-be/entity"
	"base-be/repository" // Import interface repository
	"errors"
)

// --- KONTRAK (INTERFACE) ---

// BookService adalah interface (kontrak) untuk service buku.
// Handler akan bergantung pada interface ini.
type BookService interface {
	Create(bookRequest dto.BookRequest) (entity.Book, error)
	FindAll() ([]entity.Book, error)
	FindByID(ID int) (entity.Book, error)
	Update(ID int, bookRequest dto.BookRequest) (entity.Book, error)
	Delete(ID int) error
}

// --- IMPLEMENTASI ---

// bookService adalah implementasi dari BookService
type bookService struct {
	// 'bookService' butuh 'repository' untuk berbicara ke DB
	bookRepository repository.BookRepository // Dependensinya adalah interface
}

// NewBookService adalah constructor untuk 'bookService'.
// Ini akan dipanggil di cmd/api/main.go (saat wiring)
// Perhatikan: dia menerima interface, dan mengembalikan interface
func NewBookService(r repository.BookRepository) BookService {
	return &bookService{
		bookRepository: r,
	}
}

// Implementasi method Create
func (s *bookService) Create(bookRequest dto.BookRequest) (entity.Book, error) {
	// 1. Logika Bisnis: Konversi DTO (input) ke Entity (database model)
	book := entity.Book{
		Title:       bookRequest.Title,
		Description: bookRequest.Description,
		Price:       bookRequest.Price,
		Rating:      bookRequest.Rating,
	}

	// 2. Panggil Repository untuk menyimpan ke DB
	newBook, err := s.bookRepository.Create(book)
	return newBook, err
}

// Implementasi method FindAll
func (s *bookService) FindAll() ([]entity.Book, error) {
	// (Contoh logika bisnis: bisa ditambahkan filter/pagination di sini)
	books, err := s.bookRepository.FindAll()
	return books, err
}

// Implementasi method FindByID
func (s *bookService) FindByID(ID int) (entity.Book, error) {
	book, err := s.bookRepository.FindByID(ID)
	// Logika bisnis kecil: custom error message
	if err != nil {
		return book, errors.New("buku tidak ditemukan")
	}
	return book, nil
}

// Implementasi method Update
func (s *bookService) Update(ID int, bookRequest dto.BookRequest) (entity.Book, error) {
	// 1. Logika Bisnis: Cek dulu apakah buku ada
	book, err := s.bookRepository.FindByID(ID)
	if err != nil {
		return book, errors.New("buku tidak ditemukan")
	}

	// 2. Update field-fieldnya dari DTO
	book.Title = bookRequest.Title
	book.Description = bookRequest.Description
	book.Price = bookRequest.Price
	book.Rating = bookRequest.Rating

	// 3. Panggil Repository untuk update
	updatedBook, err := s.bookRepository.Update(book)
	return updatedBook, err
}

// Implementasi method Delete
func (s *bookService) Delete(ID int) error {
	// 1. Logika Bisnis: Cek dulu apakah buku ada
	book, err := s.bookRepository.FindByID(ID)
	if err != nil {
		return errors.New("buku tidak ditemukan")
	}

	// 2. Panggil Repository untuk delete
	// Kita passing 'book' (entity) yang sudah didapat
	err = s.bookRepository.Delete(book)
	return err
}
