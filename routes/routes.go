package routes

import (
	"github.com/dragranzer/Golang-Mini-Project/factory"
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {

	_presenter_book := factory.Init()
	_presenter_author := factory.InitAuthor()

	e := echo.New()

	e.GET("/books", _presenter_book.BookPresentation.GetAllBook)
	e.POST("/books", _presenter_book.BookPresentation.InsertBook)

	e.GET("/authors", _presenter_author.AuthorPresentation.GetAllAuthor)
	e.POST("/authors", _presenter_author.AuthorPresentation.InsertAuthor)

	return e
}
