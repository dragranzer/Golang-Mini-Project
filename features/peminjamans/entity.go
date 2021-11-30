package peminjamans

type Core struct {
	ID         int    `json:"id" form:"id"`
	Hari       int    `json:"hari" form:"hari"`
	TotalHarga int    `json:"total_harga" form:"total_harga"`
	BookID     int    `json:"book_id" form:"book_id"`
	UserID     int    `json:"user_id" form:"user_id"`
	BookName   string `json:"book_name" form:"book_name"`
	UserName   string `json:"user_name" form:"user_name"`
}

type DetailBookPinjam struct {
	ID        int
	Peminjams []Peminjam
}

type Peminjam struct {
	ID   int
	Nama string
}

type Bussiness interface {
	CreateData(data Core) (resp Core, err error)
	GetAllData() (resp []Core)
	GetDetailBookPinjam(judul string) (resp DetailBookPinjam)
	// continue the function abtraction
}

type Data interface {
	InsertData(data Core) (resp Core, err error)
	SelectAllData() (resp []Core)
	SelectDetailBookPinjam(id int) (resp []Core)
}
