package books

import "time"

type Core struct {
	ID            int            `json:"id" form:"id" gorm:"primaryKey"`
	Judul         string         `json:"judul" form:"judul"`
	Tersedia      int            `json:"tersedia" form:"tersedia"`
	PublishedAt   time.Time      `json:"published_at" form:"published_at"`
	Like          int            `json:"like" form:"like"`
	Harga         int            `json:"harga" form:"harga"`
	KategoriCores []KategoriCore `json:"kategori" form:"kategori" gorm:"foreignKey:BookID;references:ID"`
	Authors       []AuthorCore   `json:"author" form:"author" gorm:"many2many:detail_book;"`
	Peminjamen    []peminjaman   `json:"peminjaman" form:"peminjaman" gorm:"foreignKey:BookID;references:ID"`
	CreatedAt     time.Time      `json:"created_at" form:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at" form:"updated_at"`
}

type peminjaman struct {
	ID         int `json:"id" form:"id" gorm:"primaryKey"`
	Hari       int `json:"hari" form:"hari"`
	TotalHarga int `json:"total_harga" form:"total_harga"`
	BookID     int `json:"book_id" form:"book_id"`
}

type KategoriCore struct {
	ID        int       `json:"id" form:"id"`
	Nama      string    `json:"nama" form:"nama"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
	BookID    int       `json:"book_id" form:"book_id"`
}

type AuthorCore struct {
	ID        int       `json:"id" form:"id"`
	Nama      string    `json:"nama" form:"nama"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
}

type Bussiness interface {
	CreateData(data Core) (resp Core, err error)
	GetAllData() (resp []Core)
	GetDetailData(name string) (resp []Core, err error)
	ChangeTersediabyName(name string) (resp Core, err error)
	// continue the function abtraction
}

type Data interface {
	InsertData(data Core) (resp Core, err error)
	SelectAllData() (resp []Core)
	SelectData(name string) (resp []Core, err error)
	UpdateTersediabyName(name string) (resp Core, err error)
}
