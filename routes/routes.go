package routes

import (
	"github.com/dragranzer/Golang-Mini-Project/factory"
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {

	_presenter_book_author_user_detbok_peminjaman := factory.Init()
	// _presenter_peminjaman := factory.InitPeminjaman()

	e := echo.New()

	// routes := e.Group("/books")

	e.GET("/books/:judul", _presenter_book_author_user_detbok_peminjaman.BookPresentation.GetBook)
	// e.GET("/books", _presenter_book.BookPresentation.GetBook)
	e.GET("/books", _presenter_book_author_user_detbok_peminjaman.BookPresentation.GetAllBook)
	e.POST("/books", _presenter_book_author_user_detbok_peminjaman.BookPresentation.InsertBook)

	e.GET("/authors", _presenter_book_author_user_detbok_peminjaman.AuthorPresentation.GetAllAuthor)
	e.POST("/authors", _presenter_book_author_user_detbok_peminjaman.AuthorPresentation.InsertAuthor)

	e.GET("/users", _presenter_book_author_user_detbok_peminjaman.UserPresentation.GetAllUser)
	e.GET("/users/:nama", _presenter_book_author_user_detbok_peminjaman.UserPresentation.GetUser)
	e.POST("/users", _presenter_book_author_user_detbok_peminjaman.UserPresentation.InsertUser)

	e.GET("/peminjamans", _presenter_book_author_user_detbok_peminjaman.PeminjamanPresentation.GetAllPeminjaman)
	e.POST("/peminjamans", _presenter_book_author_user_detbok_peminjaman.PeminjamanPresentation.InsertPeminjaman)

	return e
}
