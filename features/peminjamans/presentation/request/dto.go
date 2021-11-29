package request

import "github.com/dragranzer/Golang-Mini-Project/features/peminjamans"

type Peminjaman struct {
	ID         int `json:"id" form:"id"`
	Hari       int `json:"hari" form:"hari"`
	TotalHarga int `json:"total_harga" form:"total_harga"`
	BookID     int `json:"book_id" form:"book_id"`
	UserID     int `json:"user_id" form:"user_id"`
}

func ToCore(req Peminjaman) peminjamans.Core {
	return peminjamans.Core{
		Hari:       req.Hari,
		TotalHarga: req.TotalHarga,
		BookID:     req.BookID,
		UserID:     req.UserID,
	}
}
