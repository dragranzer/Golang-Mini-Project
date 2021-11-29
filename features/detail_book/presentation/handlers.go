package presentation

import (
	"net/http"

	"github.com/dragranzer/Golang-Mini-Project/features/detail_book"
	"github.com/dragranzer/Golang-Mini-Project/features/detail_book/presentation/request"
	"github.com/labstack/echo/v4"
)

type Detail_bookHandler struct {
	detailbookBussiness detail_book.Bussiness
}

func Newdetail_bookHandler(ab detail_book.Bussiness) *Detail_bookHandler {
	return &Detail_bookHandler{
		detailbookBussiness: ab,
	}
}

func (ah *Detail_bookHandler) InsertDetailBook(c echo.Context) error {
	// fmt.Println("Masuk Handlers F2")
	detail := request.Detail{}
	c.Bind(&detail)
	// fmt.Println("detail presentation ========== ", detail)
	data, err := ah.detailbookBussiness.CreateData(request.ToCore(detail))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"detail":  data,
	})
}
