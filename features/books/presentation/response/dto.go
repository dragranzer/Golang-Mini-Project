package response

import (
	"time"

	"github.com/dragranzer/Golang-Mini-Project/features/books"
)

type Book struct {
	Judul       string    `json:"judul" form:"judul"`
	Tersedia    int       `json:"tersedia" form:"tersedia"`
	PublishedAt time.Time `json:"published_at" form:"published_at"`
	Like        int       `json:"like" form:"like"`
	Harga       int       `json:"harga" form:"harga"`
}

func FromCore(req books.Core) Book {
	return Book{
		Judul:       req.Judul,
		Tersedia:    req.Tersedia,
		PublishedAt: req.PublishedAt,
		Like:        req.Like,
		Harga:       req.Harga,
	}
}

func FromCoreSlice(core []books.Core) []Book {
	var artArray []Book
	for key := range core {
		artArray = append(artArray, FromCore(core[key]))
	}
	return artArray
}
