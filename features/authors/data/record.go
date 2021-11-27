package data

import (
	"time"

	"github.com/dragranzer/Golang-Mini-Project/features/authors"
)

type Author struct {
	ID        int       `json:"id" form:"id"`
	Nama      string    `json:"nama" form:"nama" gorm:"UNIQUE_INDEX"`
	Book      []Book    `json:"book" form:"book" gorm:"many2many:detail_book;"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
}

type Book struct {
	ID        int       `json:"id" form:"id"`
	Judul     string    `json:"judul" form:"judul"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
}

//DTO

func (a *Author) toCore() authors.Core {
	return authors.Core{
		ID:        int(a.ID),
		Nama:      a.Nama,
		CreatedAt: a.CreatedAt,
		UpdatedAt: a.UpdatedAt,
	}
}

func toCoreList(resp []Author) []authors.Core {
	a := []authors.Core{}
	for key := range resp {
		a = append(a, resp[key].toCore())
	}
	return a
}

func fromCore(core authors.Core) Author {

	return Author{
		Nama:      core.Nama,
		CreatedAt: core.CreatedAt,
		UpdatedAt: core.UpdatedAt,
	}
}
