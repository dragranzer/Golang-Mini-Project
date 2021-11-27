package books

import "time"

type Core struct {
	ID          int            `json:"id" form:"id"`
	Judul       string         `json:"judul" form:"judul"`
	Tersedia    int            `json:"tersedia" form:"tersedia"`
	PublishedAt time.Time      `json:"published_at" form:"published_at"`
	Like        int            `json:"like" form:"like"`
	Harga       int            `json:"harga" form:"harga"`
	Kategoris   []KategoriCore `json:"kategori" form:"kategori"`
	Authors     []AuthorCore   `json:"author" form:"author" gorm:"many2many:detail_book;"`
	CreatedAt   time.Time      `json:"created_at" form:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at" form:"updated_at"`
}

type KategoriCore struct {
	ID        int       `json:"id" form:"id"`
	Nama      string    `json:"nama" form:"nama"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
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
	// continue the function abtraction
}

type Data interface {
	InsertData(data Core) (resp Core, err error)
	SelectAllData() (resp []Core)
}
