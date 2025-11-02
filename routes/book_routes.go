package routes

import (
	"base-be/handler" // Sesuaikan path ke handler Anda
	"github.com/gin-gonic/gin"
)

// SetupBookRoutes mendaftarkan semua rute yang berhubungan dengan buku.
// Perhatikan parameternya:
// 1. rg (*gin.RouterGroup): Ini adalah grup rute (seperti /v1) yang dikirim dari hub.
// 2. bookHandler (*handler.BookHandler): Handler yang berisi logika untuk rute ini.
func SetupBookRoutes(rg *gin.RouterGroup, bookHandler *handler.BookHandler) {

	// Buat sub-grup spesifik untuk /books
	// Hasilnya akan menjadi: /v1/books
	bookRoutes := rg.Group("/books")
	{
		// Definisikan semua rute yang terkait dengan buku di sini

		// POST /v1/books
		bookRoutes.POST("/", bookHandler.PostBookHandler)

		// GET /v1/books
		bookRoutes.GET("/", bookHandler.GetBooksHandler)

		// GET /v1/books/1
		bookRoutes.GET("/:id", bookHandler.GetBookByIDHandler)

		// PUT /v1/books/1
		bookRoutes.PUT("/:id", bookHandler.UpdateBookHandler)

		// DELETE /v1/books/1
		bookRoutes.DELETE("/:id", bookHandler.DeleteBookHandler)
	}
}
