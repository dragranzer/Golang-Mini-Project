package factory

import (
	"github.com/dragranzer/Golang-Mini-Project/config"
	_book_bussiness "github.com/dragranzer/Golang-Mini-Project/features/books/bussiness"
	_book_data "github.com/dragranzer/Golang-Mini-Project/features/books/data"
	_book_presentation "github.com/dragranzer/Golang-Mini-Project/features/books/presentation"
)

type Presenter struct {
	BookPresentation _book_presentation.BooksHandler
}

func Init() Presenter {

	bookData := _book_data.NewBookRepository(config.DB)
	bookBussiness := _book_bussiness.NewAricleBussiness(bookData)

	return Presenter{
		BookPresentation: *_book_presentation.NewBookHandler(bookBussiness),
	}
}
