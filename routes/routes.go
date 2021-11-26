package routes

import (
	"github.com/dragranzer/Golang-Mini-Project/factory"
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {

	presenter := factory.Init()

	e := echo.New()

	e.GET("/books", presenter.BookPresentation.GetAllBook)
	e.POST("/books", presenter.BookPresentation.InsertBook)

	return e
}
