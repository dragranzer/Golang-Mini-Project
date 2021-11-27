package data

import (
	"time"

	"github.com/dragranzer/Golang-Mini-Project/features/books"
)

type Book struct {
	ID          int        `json:"id" form:"id"`
	Judul       string     `json:"judul" form:"judul"`
	Tersedia    int        `json:"tersedia" form:"tersedia"`
	PublishedAt time.Time  `json:"published_at" form:"published_at"`
	Like        int        `json:"like" form:"like"`
	Harga       int        `json:"harga" form:"harga"`
	Kategori    []Kategori `json:"kategori" form:"kategori" gorm:"foreigKey:BookID;references:ID"`
	Author      []Author   `json:"author" form:"author" gorm:"many2many:detail_book;"`
	CreatedAt   time.Time  `json:"created_at" form:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at" form:"updated_at"`
}

type Kategori struct {
	ID        int       `json:"id" form:"id"`
	Nama      string    `json:"nama" form:"nama"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
	BookID    int
}

type Author struct {
	ID        int       `json:"id" form:"id"`
	Nama      string    `json:"nama" form:"nama"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
}

//DTO

func (a *Book) toCore() books.Core {
	return books.Core{
		ID:          int(a.ID),
		Judul:       a.Judul,
		Tersedia:    a.Tersedia,
		PublishedAt: a.PublishedAt,
		Like:        a.Like,
		Harga:       a.Harga,
		CreatedAt:   a.CreatedAt,
		UpdatedAt:   a.UpdatedAt,
	}
}

func toCoreList(resp []Book) []books.Core {
	a := []books.Core{}
	for key := range resp {
		a = append(a, resp[key].toCore())
	}
	return a
}

func fromCore(core books.Core) Book {

	return Book{
		Judul:       core.Judul,
		Tersedia:    core.Tersedia,
		PublishedAt: core.PublishedAt,
		Like:        core.Like,
		Harga:       core.Harga,
		CreatedAt:   core.CreatedAt,
		UpdatedAt:   core.UpdatedAt,
	}
}
