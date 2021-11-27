package authors

import "time"

type Core struct {
	ID        int        `json:"id" form:"id"`
	Nama      string     `json:"nama" form:"nama"`
	Books     []BookCore `json:"book" form:"book" gorm:"many2many:detail_book;"`
	CreatedAt time.Time  `json:"created_at" form:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" form:"updated_at"`
}

type BookCore struct {
	ID        int       `json:"id" form:"id"`
	Judul     string    `json:"judul" form:"judul"`
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
