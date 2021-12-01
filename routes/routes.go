package routes

import (
	"github.com/dragranzer/Golang-Mini-Project/factory"
	"github.com/dragranzer/Golang-Mini-Project/middleware"
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {

	_presenter := factory.Init()

	e := echo.New()
	middleware.LogMidd(e)
	// routes := e.Group("/books")

	e.GET("/books/:judul", _presenter.BookPresentation.GetBook)
	// e.GET("/books", _presenter_book.BookPresentation.GetBook)
	e.GET("/books", _presenter.BookPresentation.GetAllBook)
	e.POST("/books", _presenter.BookPresentation.InsertBook)

	e.GET("/authors", _presenter.AuthorPresentation.GetAllAuthor)
	e.POST("/authors", _presenter.AuthorPresentation.InsertAuthor)

	e.GET("/users", _presenter.UserPresentation.GetAllUser)
	e.GET("/users/:nama", _presenter.UserPresentation.GetUser)
	e.POST("/users", _presenter.UserPresentation.InsertUser)
	e.PUT("/users/:id", _presenter.UserPresentation.UpdateUserData)
	e.GET("/login", _presenter.UserPresentation.LoginUser)

	e.GET("/peminjamans", _presenter.PeminjamanPresentation.GetAllPeminjaman)
	e.POST("/peminjamans", _presenter.PeminjamanPresentation.InsertPeminjaman)
	e.GET("/peminjamans/book/:judul", _presenter.PeminjamanPresentation.GetDetailPinjam)

	e.POST("/favorite", _presenter.FavoritePresentation.InsertFavorites)
	e.GET("/favoriteofuser/:id", _presenter.FavoritePresentation.GetFavbyUserID)
	e.GET("/favoriteofbook/:id", _presenter.FavoritePresentation.GetFavbyBookID)

	return e
}
