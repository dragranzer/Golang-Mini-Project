package peminjamans

type Core struct {
	ID         int `json:"id" form:"id"`
	Hari       int `json:"hari" form:"hari"`
	TotalHarga int `json:"total_harga" form:"total_harga"`
	BookID     int `json:"book_id" form:"book_id"`
	UserID     int `json:"user_id" form:"user_id"`
}

type Bussiness interface {
	CreateData(data Core) (resp Core, err error)
	GetAllData() (resp []Core)
	// continue the function abtraction
}

type Data interface {
	InsertData(data Core) (resp Core, err error)
	SelectAllData() (resp []Core)
}
