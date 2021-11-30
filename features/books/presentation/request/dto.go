package request

import (
	"time"

	"github.com/dragranzer/Golang-Mini-Project/features/books"
)

type Book struct {
	Judul       string       `json:"judul" form:"judul"`
	Tersedia    int          `json:"tersedia" form:"tersedia"`
	PublishedAt time.Time    `json:"published_at" form:"published_at"`
	Like        int          `json:"like" form:"like"`
	Harga       int          `json:"harga" form:"harga"`
	Authors     []AuthorCore `json:"author" form:"author"`
}

type AuthorCore struct {
	Nama string `json:"nama" form:"nama"`
}

func ToCore(req Book) books.Core {
	AuthorCore := []books.AuthorCore{}
	for _, value := range req.Authors {
		AuthorCore = append(AuthorCore, books.AuthorCore{
			Nama: value.Nama,
		})
	}
	// fmt.Println("list Author", AuthorCore)
	return books.Core{
		Judul:       req.Judul,
		Tersedia:    req.Tersedia,
		PublishedAt: req.PublishedAt,
		Like:        req.Like,
		Harga:       req.Harga,
		Authors:     AuthorCore,
	}
}
