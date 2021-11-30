package presentation

import (
	"net/http"

	"github.com/dragranzer/Golang-Mini-Project/features/favorites"
	"github.com/dragranzer/Golang-Mini-Project/features/favorites/presentation/request"
	"github.com/dragranzer/Golang-Mini-Project/features/favorites/presentation/response"
	"github.com/labstack/echo/v4"
)

type FavoritesHandler struct {
	favoriteBussiness favorites.Bussiness
}

func NewFavoriteHandler(ab favorites.Bussiness) *FavoritesHandler {
	return &FavoritesHandler{
		favoriteBussiness: ab,
	}
}

func (ah *FavoritesHandler) InsertFavorites(c echo.Context) error {
	fav := request.Favorite{}
	c.Bind(&fav)
	err := ah.favoriteBussiness.CreateData(request.ToCore(fav))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success melakukan insert",
	})
}

func (ah *FavoritesHandler) GetFavbyUserID(c echo.Context) error {
	fav := request.Favorite{}
	c.Bind(&fav)
	resp, err := ah.favoriteBussiness.GetDatabyUserID(fav.UserID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "ini dia datanya :)",
		"data":    response.FromCore(resp),
	})
}
