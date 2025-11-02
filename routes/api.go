package routes

import (
	"base-be/handler" // Sesuaikan path ke handler Anda
	"github.com/gin-gonic/gin"
)

// SetupRoutes adalah "hub" utama untuk mendaftarkan SEMUA rute di aplikasi.
// main.go akan memanggil fungsi ini.
func SetupRoutes(r *gin.Engine,
	bookHandler *handler.BookHandler,
// userHandler *handler.UserHandler, // <--- (Contoh jika nanti Anda punya modul user)
) {

	// Rute publik / non-versioned (Contoh: health check)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// --- Grup Rute API v1 ---
	v1 := r.Group("/v1")
	{
		// Delegasikan pendaftaran rute buku ke fungsi spesifiknya
		SetupBookRoutes(v1, bookHandler)

		// (Nanti, jika Anda punya modul user)
		// SetupUserRoutes(v1, userHandler)

		// (Nanti, jika Anda punya modul order)
		// SetupOrderRoutes(v1, orderHandler)
	}

	// --- Grup Rute API v2 (Jika ada di masa depan) ---
	// v2 := r.Group("/v2")
	// {
	//    ...
	// }
}
