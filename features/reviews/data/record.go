package data

import "github.com/dragranzer/Golang-Mini-Project/features/reviews"

type Review struct {
	UserID int
	BookID int
	Review string
}

//DTO

func (a *Review) toCore() reviews.Core {
	return reviews.Core{
		BookID: a.BookID,
		UserID: a.UserID,
		Review: a.Review,
	}
}

func toCoreList(resp []Review) []reviews.Core {
	a := []reviews.Core{}
	for key := range resp {
		a = append(a, resp[key].toCore())
	}
	return a
}

func fromCore(core reviews.Core) Review {

	return Review{
		BookID: core.BookID,
		UserID: core.UserID,
		Review: core.Review,
	}
}
