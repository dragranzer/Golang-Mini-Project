package routes

import (
	"github.com/dragranzer/Golang-Mini-Project/factory"
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {

	_presenter_book := factory.Init()
	_presenter_author := factory.InitAuthor()
	_presenter_user := factory.InitUser()
	_presenter_peminjaman := factory.InitPeminjaman()

	e := echo.New()

	// routes := e.Group("/books")

	e.GET("/books/:judul", _presenter_book.BookPresentation.GetBook)
	// e.GET("/books", _presenter_book.BookPresentation.GetBook)
	e.GET("/books", _presenter_book.BookPresentation.GetAllBook)
	e.POST("/books", _presenter_book.BookPresentation.InsertBook)

	e.GET("/authors", _presenter_author.AuthorPresentation.GetAllAuthor)
	e.POST("/authors", _presenter_author.AuthorPresentation.InsertAuthor)

	e.GET("/users", _presenter_user.UserPresentation.GetAllUser)
	e.POST("/users", _presenter_user.UserPresentation.InsertUser)

	e.GET("/peminjamans", _presenter_peminjaman.PeminjamanPresentation.GetAllPeminjaman)
	e.POST("/peminjamans", _presenter_peminjaman.PeminjamanPresentation.InsertPeminjaman)

	return e
}
