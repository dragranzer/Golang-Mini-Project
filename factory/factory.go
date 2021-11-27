package factory

import (
	"github.com/dragranzer/Golang-Mini-Project/config"
	_book_bussiness "github.com/dragranzer/Golang-Mini-Project/features/books/bussiness"
	_book_data "github.com/dragranzer/Golang-Mini-Project/features/books/data"
	_book_presentation "github.com/dragranzer/Golang-Mini-Project/features/books/presentation"
	// _author_bussiness "github.com/dragranzer/Golang-Mini-Project/features/authors/bussiness"
	// _author_data "github.com/dragranzer/Golang-Mini-Project/features/authors/data"
	// _author_presentation "github.com/dragranzer/Golang-Mini-Project/features/authors/presentation"
)

type Presenter struct {
	BookPresentation _book_presentation.BooksHandler
	// AuthorPresentation _author_presentation.AuthorsHandler
}

func Init() Presenter {

	bookData := _book_data.NewBookRepository(config.DB)
	bookBussiness := _book_bussiness.NewBookBussiness(bookData)

	// authorData := _author_data.NewAuthorRepository(config.DB)
	// authorBussiness := _author_bussiness.NewAuthorBussiness(authorData)

	return Presenter{
		BookPresentation: *_book_presentation.NewBookHandler(bookBussiness),
		// AuthorPresentation: *_author_presentation.NewAuthorHandler(authorBussiness),
	}
}
