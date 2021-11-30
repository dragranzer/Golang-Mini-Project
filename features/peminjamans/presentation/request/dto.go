package request

import "github.com/dragranzer/Golang-Mini-Project/features/peminjamans"

type Peminjaman struct {
	Hari       int    `json:"hari" form:"hari"`
	TotalHarga int    `json:"total_harga" form:"total_harga"`
	BookName   string `json:"book_name" form:"book_name"`
	UserName   string `json:"user_name" form:"user_name"`
}

func ToCore(req Peminjaman) peminjamans.Core {
	return peminjamans.Core{
		Hari:       req.Hari,
		TotalHarga: req.TotalHarga,
		BookName:   req.BookName,
		UserName:   req.UserName,
	}
}
