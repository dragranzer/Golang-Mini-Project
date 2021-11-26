package presentation

import (
	"fmt"
	"net/http"

	"github.com/dragranzer/Golang-Mini-Project/features/books"
	"github.com/dragranzer/Golang-Mini-Project/features/books/presentation/request"
	"github.com/dragranzer/Golang-Mini-Project/features/books/presentation/response"
	"github.com/labstack/echo/v4"
)

type BooksHandler struct {
	bookBussiness books.Bussiness
}

func NewBookHandler(ab books.Bussiness) *BooksHandler {
	return &BooksHandler{
		bookBussiness: ab,
	}
}

func (boh *BooksHandler) GetAllBook(c echo.Context) error {
	fmt.Println("Masuk Handlers F2")
	result := boh.bookBussiness.GetAllData()

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "hope all feeling well",
		"data":    response.FromCoreSlice(result),
	})
}

func (ah *BooksHandler) InsertBook(c echo.Context) error {
	// fmt.Println("Masuk Handlers F2")
	book := request.Book{}
	c.Bind(&book)
	fmt.Println("book presentation ========== ", book)
	data, err := ah.bookBussiness.CreateData(request.ToCore(book))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"book":    data,
	})
}
