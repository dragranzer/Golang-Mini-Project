package request

import "github.com/dragranzer/Golang-Mini-Project/features/reviews"

type Review struct {
	BookID int    `json:"book_id" form:"book_id"`
	UserID int    `json:"user_id" form:"user_id"`
	Review string `json:"review" form:"review"`
}

func ToCore(req Review) reviews.Core {
	return reviews.Core{
		BookID: req.BookID,
		UserID: req.UserID,
		Review: req.Review,
	}
}
