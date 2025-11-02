package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"base-be/dto"     // Kita akan buat file DTO ini nanti
	"base-be/service" // Kita akan buat service ini nanti

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// BookHandler adalah struct yang akan menampung 'service' sebagai dependency
type BookHandler struct {
	// Dependency-nya adalah 'service', tapi dalam bentuk INTERFACE
	// Ini membuat handler kita 'testable' dan 'decoupled' (terpisah)
	bookService service.BookService
}

// NewBookHandler adalah 'constructor' untuk BookHandler
// Ini akan dipanggil di main.go (wiring)
func NewBookHandler(s service.BookService) *BookHandler {
	return &BookHandler{
		bookService: s,
	}
}

// --- METODE-METODE HANDLER ---
// Nama-nama method ini HARUS SESUAI dengan yang kita panggil di 'routes/book_routes.go'

// PostBookHandler menangani request POST /v1/books
func (h *BookHandler) PostBookHandler(c *gin.Context) {
	// 1. Siapkan variabel untuk menampung input JSON
	var bookRequest dto.BookRequest

	// 2. Bind JSON ke struct 'bookRequest'
	// ShouldBindJSON juga otomatis memvalidasi tag 'binding' di DTO
	err := c.ShouldBindJSON(&bookRequest)

	// 3. Handle jika ada error binding atau validasi
	if err != nil {
		// Kita akan format error validasi agar lebih cantik
		errorMessages := []string{}
		var validationErrs validator.ValidationErrors

		if errors.As(err, &validationErrs) {
			for _, e := range validationErrs {
				// Format pesan error: "Error pada field 'Nama', kondisi: 'required'"
				errorMessage := fmt.Sprintf("Error pada field %s, kondisi: %s", e.Field(), e.ActualTag())
				errorMessages = append(errorMessages, errorMessage)
			}
			c.JSON(http.StatusBadRequest, gin.H{
				"errors": errorMessages,
			})
			return
		}

		// Handle error jika JSON-nya tidak valid (bukan error validasi)
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}

	// 4. Panggil Service
	// Jika tidak ada error, 'bookRequest' sudah terisi data
	// Kita lempar ke service layer
	newBook, err := h.bookService.Create(bookRequest)
	if err != nil {
		// Jika service mengembalikan error (misal: gagal simpan)
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}

	// 5. Kembalikan respons sukses
	c.JSON(http.StatusCreated, gin.H{ // 201 Created lebih cocok untuk POST
		"data": newBook,
	})
}

// GetBooksHandler menangani request GET /v1/books
func (h *BookHandler) GetBooksHandler(c *gin.Context) {
	// 1. Panggil Service untuk ambil semua buku
	books, err := h.bookService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}

	// 2. Kembalikan respons sukses
	c.JSON(http.StatusOK, gin.H{
		"data": books,
	})
}

// GetBookByIDHandler menangani request GET /v1/books/:id
func (h *BookHandler) GetBookByIDHandler(c *gin.Context) {
	// 1. Ambil 'id' dari parameter URL
	idStr := c.Param("id")
	// 2. Konversi ID ke integer
	id, err := strconv.Atoi(idStr)
	if err != nil {
		// Jika ID bukan angka
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": "ID tidak valid",
		})
		return
	}

	// 3. Panggil Service
	book, err := h.bookService.FindByID(id)
	if err != nil {
		// Jika service mengembalikan error (misal: data not found)
		c.JSON(http.StatusNotFound, gin.H{
			"errors": err.Error(),
		})
		return
	}

	// 4. Kembalikan respons sukses
	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}

// UpdateBookHandler menangani request PUT /v1/books/:id
func (h *BookHandler) UpdateBookHandler(c *gin.Context) {
	// 1. Ambil ID dari URL dan konversi ke int (sama seperti GetBookByIDHandler)
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": "ID tidak valid"})
		return
	}

	// 2. Bind data JSON yang baru (sama seperti PostBookHandler)
	var bookRequest dto.BookRequest
	err = c.ShouldBindJSON(&bookRequest)
	if err != nil {
		// (Anda bisa pakai error handling validasi yang detail seperti di PostBookHandler)
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}

	// 3. Panggil Service
	updatedBook, err := h.bookService.Update(id, bookRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}

	// 4. Kembalikan respons sukses
	c.JSON(http.StatusOK, gin.H{
		"data": updatedBook,
	})
}

// DeleteBookHandler menangani request DELETE /v1/books/:id
func (h *BookHandler) DeleteBookHandler(c *gin.Context) {
	// 1. Ambil ID dari URL dan konversi ke int
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": "ID tidak valid"})
		return
	}

	// 2. Panggil Service
	err = h.bookService.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}

	// 3. Kembalikan respons sukses
	c.JSON(http.StatusOK, gin.H{
		"message": "Buku berhasil dihapus",
	})
}
