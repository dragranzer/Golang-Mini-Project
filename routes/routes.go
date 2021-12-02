package routes

import (
	"github.com/dragranzer/Golang-Mini-Project/constant"
	"github.com/dragranzer/Golang-Mini-Project/factory"
	"github.com/dragranzer/Golang-Mini-Project/middleware"
	"github.com/labstack/echo/v4"
	mid "github.com/labstack/echo/v4/middleware"
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

	eJWT := e.Group("")
	eJWT.Use(mid.JWT([]byte(constant.SECRET_JWT)))

	eJWT.GET("/users", _presenter.UserPresentation.GetAllUser)
	e.GET("/users/:nama", _presenter.UserPresentation.GetUser)
	e.POST("/users", _presenter.UserPresentation.InsertUser)
	e.PUT("/users/:id", _presenter.UserPresentation.UpdateUserData)
	e.GET("/login", _presenter.UserPresentation.LoginUser)

	eJWT.GET("/peminjamans", _presenter.PeminjamanPresentation.GetAllPeminjaman)
	eJWT.POST("/peminjamans", _presenter.PeminjamanPresentation.InsertPeminjaman)
	eJWT.GET("/peminjamans/book/:judul", _presenter.PeminjamanPresentation.GetDetailPinjam)

	eJWT.POST("/favorite", _presenter.FavoritePresentation.InsertFavorites)
	eJWT.DELETE("/favorite", _presenter.FavoritePresentation.DeteleFavbyBookID)
	eJWT.GET("/favoriteofuser/:id", _presenter.FavoritePresentation.GetFavbyUserID)
	eJWT.GET("/favoriteofbook/:id", _presenter.FavoritePresentation.GetFavbyBookID)

	e.GET("/reviews", _presenter.ReviewPresentation.GetFavbyBookID)
	e.POST("/reviews", _presenter.ReviewPresentation.InsertReviews)
	e.DELETE("/reviews", _presenter.ReviewPresentation.DeletebyBookID)

	return e
}
