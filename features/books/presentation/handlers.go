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
	fmt.Println("book0 = ", book)
	c.Bind(&book)
	fmt.Println("book = ", book)
	fmt.Println("book presentation ========== ", book)
	_, err := ah.bookBussiness.CreateData(request.ToCore(book))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success melakukan insert",
	})
}

func (boh *BooksHandler) GetBook(c echo.Context) error {
	// fmt.Println("masuk handlers")
	var judul string
	echo.PathParamsBinder(c).String("judul", &judul)
	// fmt.Println("judul = ", judul)
	result, err := boh.bookBussiness.GetDetailData(judul)

	if err != nil {
		return c.JSON(404, map[string]interface{}{
			"message": "Could not get article",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "hope all feeling well",
		"data":    response.FromCoreSlice(result),
	})
}
