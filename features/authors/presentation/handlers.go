package presentation

import (
	"fmt"
	"net/http"

	"github.com/dragranzer/Golang-Mini-Project/features/authors"
	"github.com/dragranzer/Golang-Mini-Project/features/authors/presentation/request"
	"github.com/dragranzer/Golang-Mini-Project/features/authors/presentation/response"
	"github.com/labstack/echo/v4"
)

type AuthorsHandler struct {
	authorBussiness authors.Bussiness
}

func NewAuthorHandler(ab authors.Bussiness) *AuthorsHandler {
	return &AuthorsHandler{
		authorBussiness: ab,
	}
}

func (boh *AuthorsHandler) GetAllAuthor(c echo.Context) error {
	fmt.Println("Masuk Handlers F2")
	result := boh.authorBussiness.GetAllData()

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "hope all feeling well",
		"data":    response.FromCoreSlice(result),
	})
}

func (ah *AuthorsHandler) InsertAuthor(c echo.Context) error {
	// fmt.Println("Masuk Handlers F2")
	author := request.Author{}
	c.Bind(&author)
	// fmt.Println("author presentation ========== ", author)
	data, err := ah.authorBussiness.CreateData(request.ToCore(author))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"author":  data,
	})
}
