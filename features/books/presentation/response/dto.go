package response

import (
	"time"

	"github.com/dragranzer/Golang-Mini-Project/features/books"
)

type Book struct {
	Judul       string     `json:"judul" form:"judul"`
	Tersedia    int        `json:"tersedia" form:"tersedia"`
	PublishedAt time.Time  `json:"published_at" form:"published_at"`
	Like        int        `json:"like" form:"like"`
	Harga       int        `json:"harga" form:"harga"`
	Kategoris   []Kategori `json:"kategori" form:"kategori"`
	Authors     []Author
}

type Kategori struct {
	ID   int    `json:"id" form:"id"`
	Nama string `json:"nama" form:"nama"`
}

type Author struct {
	ID   int    `json:"id" form:"id"`
	Nama string `json:"nama" form:"nama"`
}

func FromCore(req books.Core) Book {
	listAuthor := []Author{}
	for _, value := range req.Authors {
		listAuthor = append(listAuthor, Author{
			ID:   value.ID,
			Nama: value.Nama,
		})
	}
	return Book{
		Judul:       req.Judul,
		Tersedia:    req.Tersedia,
		PublishedAt: req.PublishedAt,
		Like:        req.Like,
		Harga:       req.Harga,
		Authors:     listAuthor,
	}
}

func FromCoreSlice(core []books.Core) []Book {
	var artArray []Book
	for key := range core {
		artArray = append(artArray, FromCore(core[key]))
	}
	return artArray
}
