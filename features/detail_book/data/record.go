package data

import "github.com/dragranzer/Golang-Mini-Project/features/detail_book"

type DetailBook struct {
	BookID   int `gorm:"primaryKey"`
	AuthorID int `gorm:"primaryKey"`
}

//DTO

// func (a *DetailBook) toCore() detail_book.Core {
// 	return detail_book.Core{
// 		BookID:   a.BookID,
// 		AuthorID: a.AuthorID,
// 	}
// }

// func toCoreList(resp []DetailBook) []detail_book.Core {
// 	a := []detail_book.Core{}
// 	for key := range resp {
// 		a = append(a, resp[key].toCore())
// 	}
// 	return a
// }

func fromCore(core detail_book.Core) DetailBook {

	return DetailBook{
		BookID:   core.BookID,
		AuthorID: core.AuthorID,
	}
}

func toCore1(book DetailBook) detail_book.Core {

	return detail_book.Core{
		BookID:   book.BookID,
		AuthorID: book.AuthorID,
	}
}
