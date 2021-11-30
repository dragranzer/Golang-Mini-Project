package response

import "github.com/dragranzer/Golang-Mini-Project/features/favorites"

type Favorite struct {
	UserID        int
	BookID        int
	FavoriteBooks []string
	FavoriteUsers []string
}

type FavoritebyUser struct {
	UserID        int
	FavoriteBooks []string
}

type FavoritebyBook struct {
	BookID        int
	FavoriteUsers []string
}

func FromCore(res favorites.Core) FavoritebyUser {
	return FavoritebyUser{
		UserID:        res.UserID,
		FavoriteBooks: res.BookName,
	}
}

func FromCoreUser(res favorites.Core) FavoritebyBook {
	return FavoritebyBook{
		BookID:        res.BookID,
		FavoriteUsers: res.UserName,
	}
}

func FromCoreSlice(core []favorites.Core) []FavoritebyUser {
	var artArray []FavoritebyUser
	for key := range core {
		artArray = append(artArray, FromCore(core[key]))
	}
	return artArray
}
