package response

import "github.com/dragranzer/Golang-Mini-Project/features/peminjamans"

type Peminjaman struct {
	Hari       int `json:"hari" form:"hari"`
	TotalHarga int `json:"total_harga" form:"total_harga"`
	BookID     int `json:"book_id" form:"book_id"`
	UserID     int `json:"user_id" form:"user_id"`
}

func FromCore(req peminjamans.Core) Peminjaman {
	return Peminjaman{
		Hari:       req.Hari,
		TotalHarga: req.TotalHarga,
		BookID:     req.BookID,
		UserID:     req.UserID,
	}
}

func FromCoreSlice(core []peminjamans.Core) []Peminjaman {
	var artArray []Peminjaman
	for key := range core {
		artArray = append(artArray, FromCore(core[key]))
	}
	return artArray
}
