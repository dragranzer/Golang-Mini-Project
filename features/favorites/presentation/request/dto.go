package request

import "github.com/dragranzer/Golang-Mini-Project/features/favorites"

type Favorite struct {
	BookID int `json:"book_id" form:"book_id"`
	UserID int `json:"user_id" form:"user_id"`
}

func ToCore(req Favorite) favorites.Core {
	return favorites.Core{
		BookID: req.BookID,
		UserID: req.UserID,
	}
}
