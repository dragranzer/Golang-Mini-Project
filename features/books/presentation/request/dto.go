package request

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

func ToCore(req Book) books.Core {
	return books.Core{
		Judul:       req.Judul,
		Tersedia:    req.Tersedia,
		PublishedAt: req.PublishedAt,
		Like:        req.Like,
		Harga:       req.Harga,
	}
}
