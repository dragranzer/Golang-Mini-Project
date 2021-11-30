package request

import "github.com/dragranzer/Golang-Mini-Project/features/detail_book"

type Detail struct {
	BookID   int `json:"book_id" form:"book_id"`
	AuthorID int `json:"author_id" form:"author_id"`
}

func ToCore(req Detail) detail_book.Core {

	return detail_book.Core{
		BookID:   req.BookID,
		AuthorID: req.AuthorID,
	}
}
