package presentation

import (
	"net/http"

	"github.com/dragranzer/Golang-Mini-Project/features/reviews"
	"github.com/dragranzer/Golang-Mini-Project/features/reviews/presentation/request"
	"github.com/dragranzer/Golang-Mini-Project/features/reviews/presentation/response"
	"github.com/labstack/echo/v4"
)

type ReviewsHandler struct {
	reviewBussiness reviews.Bussiness
}

func NewReviewHandler(ab reviews.Bussiness) *ReviewsHandler {
	return &ReviewsHandler{
		reviewBussiness: ab,
	}
}

func (ah *ReviewsHandler) InsertReviews(c echo.Context) error {
	rev := request.Review{}
	c.Bind(&rev)
	err := ah.reviewBussiness.CreateData(request.ToCore(rev))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success melakukan insert",
	})
}

func (ah *ReviewsHandler) GetFavbyBookID(c echo.Context) error {
	fav := request.Review{}
	c.Bind(&fav)
	// fmt.Println(fav.BookID)
	resp, err := ah.reviewBussiness.GetDatabyBookID(fav.BookID)
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

func (ah *ReviewsHandler) DeletebyBookID(c echo.Context) error {
	fav := request.Review{}
	c.Bind(&fav)
	// fmt.Println(fav.BookID)
	err := ah.reviewBussiness.ClearDatabyBookID(fav.BookID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Delete data berhasil",
	})
}
