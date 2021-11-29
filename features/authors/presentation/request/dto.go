package request

import "github.com/dragranzer/Golang-Mini-Project/features/authors"

type Author struct {
	Nama  string `json:"nama" form:"nama"`
	Books []Book `json:"book" form:"book"`
}

type Book struct {
	Judul string `json:"judul" form:"judul"`
}

func ToCore(req Author) authors.Core {
	BookCore := []authors.BookCore{}
	for _, value := range req.Books {
		BookCore = append(BookCore, authors.BookCore{
			Judul: value.Judul,
		})
	}
	return authors.Core{
		Nama:  req.Nama,
		Books: BookCore,
	}
}
