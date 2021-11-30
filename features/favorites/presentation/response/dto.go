package response

import "github.com/dragranzer/Golang-Mini-Project/features/favorites"

type Favorite struct {
	UserID        int
	FavoriteBooks []string
}

func FromCore(res favorites.Core) Favorite {
	return Favorite{
		UserID:        res.UserID,
		FavoriteBooks: res.BookName,
	}
}

func FromCoreSlice(core []favorites.Core) []Favorite {
	var artArray []Favorite
	for key := range core {
		artArray = append(artArray, FromCore(core[key]))
	}
	return artArray
}
