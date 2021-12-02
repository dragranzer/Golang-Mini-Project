package response

import "github.com/dragranzer/Golang-Mini-Project/features/reviews"

type ReviewDet struct {
	BookID   int
	UserRevs []UserRev
}

type UserRev struct {
	UserID int
	Review string
}

func FromCore(res reviews.Core) ReviewDet {
	komenDetail := []UserRev{}
	for _, value := range res.ListReview {
		komenDetail = append(komenDetail, UserRev{
			UserID: value.UserID,
			Review: value.Review,
		})
	}
	return ReviewDet{
		BookID:   res.BookID,
		UserRevs: komenDetail,
	}
}
