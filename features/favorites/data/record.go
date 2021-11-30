package data

import "github.com/dragranzer/Golang-Mini-Project/features/favorites"

type Favorite struct {
	BookID int
	UserID int
}

//DTO

func (a *Favorite) toCore() favorites.Core {
	return favorites.Core{
		BookID: a.BookID,
		UserID: a.UserID,
	}
}

func toCoreList(resp []Favorite) []favorites.Core {
	a := []favorites.Core{}
	for key := range resp {
		a = append(a, resp[key].toCore())
	}
	return a
}

func fromCore(core favorites.Core) Favorite {

	return Favorite{
		BookID: core.BookID,
		UserID: core.UserID,
	}
}
