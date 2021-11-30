package bussiness

import (
	"github.com/dragranzer/Golang-Mini-Project/features/books"
	"github.com/dragranzer/Golang-Mini-Project/features/favorites"
	"github.com/dragranzer/Golang-Mini-Project/features/users"
)

type favoritesUsecase struct {
	favoriteData  favorites.Data
	bookBussiness books.Bussiness
	userBussiness users.Bussiness
}

func NewFavoriteBussiness(fD favorites.Data, bB books.Bussiness, uB users.Bussiness) favorites.Bussiness {
	return &favoritesUsecase{
		favoriteData:  fD,
		bookBussiness: bB,
		userBussiness: uB,
	}
}

func (bu *favoritesUsecase) CreateData(data favorites.Core) (err error) {
	err = bu.favoriteData.InsertData(data)
	return
}

func (bu *favoritesUsecase) GetDatabyUserID(id int) (resp favorites.Core, err error) {
	listFav, _ := bu.favoriteData.SelectDatabyUserID(id)
	listBookName := []string{}
	for _, value := range listFav {
		tempBook, _ := bu.bookBussiness.GetDetailDatabyID(value.BookID)
		listBookName = append(listBookName, tempBook.Judul)
	}
	favDetail := favorites.Core{
		UserID:   id,
		BookName: listBookName,
	}

	return favDetail, nil
}
