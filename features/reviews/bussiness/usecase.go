package bussiness

import (
	"github.com/dragranzer/Golang-Mini-Project/features/books"
	"github.com/dragranzer/Golang-Mini-Project/features/reviews"
	"github.com/dragranzer/Golang-Mini-Project/features/users"
)

type reviewsUsecase struct {
	reviewData    reviews.Data
	bookBussiness books.Bussiness
	userBussiness users.Bussiness
}

func NewReviewBussiness(fD reviews.Data, bB books.Bussiness, uB users.Bussiness) reviews.Bussiness {
	return &reviewsUsecase{
		reviewData:    fD,
		bookBussiness: bB,
		userBussiness: uB,
	}
}

func (bu *reviewsUsecase) CreateData(data reviews.Core) (err error) {
	err = bu.reviewData.InsertData(data)
	return
}

func (bu *reviewsUsecase) GetDatabyBookID(id int) (resp reviews.Core, err error) {
	listRev, _ := bu.reviewData.SelectDatabyBookID(id)
	listKomentar := []reviews.Core{}
	for _, value := range listRev {
		// tempUser, _ := bu.userBussiness.GetDatabyID(value.UserID)
		// fmt.Println(tempUser)
		listKomentar = append(listKomentar, reviews.Core{
			UserID: value.UserID,
			Review: value.Review,
		})
	}

	userAndReview := []reviews.UserRev{}
	for _, value := range listKomentar {
		userAndReview = append(userAndReview, reviews.UserRev{
			UserID: value.UserID,
			Review: value.Review,
		})
	}

	komentarOfBook := reviews.Core{
		BookID:     id,
		ListReview: userAndReview,
	}

	return komentarOfBook, nil
}

func (bu *reviewsUsecase) ClearDatabyBookID(id int) (err error) {
	err = bu.reviewData.DeleteDatabyBookID(id)
	return
}
