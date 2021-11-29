package data

import "github.com/dragranzer/Golang-Mini-Project/features/peminjamans"

type Peminjaman struct {
	ID         int `json:"id" form:"id"`
	Hari       int `json:"hari" form:"hari"`
	TotalHarga int `json:"total_harga" form:"total_harga"`
	BookID     int `json:"book_id" form:"book_id"`
	UserID     int `json:"user_id" form:"user_id"`
}

//DTO

func (a *Peminjaman) toCore() peminjamans.Core {
	return peminjamans.Core{
		ID:         int(a.ID),
		Hari:       a.Hari,
		TotalHarga: a.TotalHarga,
		BookID:     a.BookID,
		UserID:     a.UserID,
	}
}

func toCoreList(resp []Peminjaman) []peminjamans.Core {
	a := []peminjamans.Core{}
	for key := range resp {
		a = append(a, resp[key].toCore())
	}
	return a
}

func fromCore(core peminjamans.Core) Peminjaman {

	return Peminjaman{
		Hari:       core.Hari,
		TotalHarga: core.TotalHarga,
		BookID:     core.BookID,
		UserID:     core.UserID,
	}
}
