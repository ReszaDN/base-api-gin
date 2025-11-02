package dto

// BookRequest adalah DTO (Data Transfer Object) untuk 'create' dan 'update' buku.
// Di sinilah kita menaruh validasi input JSON.
// Handler akan menggunakan struct ini untuk c.ShouldBindJSON().
type BookRequest struct {
	// json:"title"      -> nama field di JSON
	// binding:"required" -> aturan validasi dari 'validator'
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Price       int    `json:"price" binding:"required,number"`
	Rating      int    `json:"rating" binding:"required,number,min=1,max=5"` // Contoh validasi tambahan
}

// Catatan:
// Kita tidak perlu BookResponse DTO terpisah jika respons JSON-nya
// sama dengan struct entity.Book. Kita bisa langsung kembalikan entity.Book.
