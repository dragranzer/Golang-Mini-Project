package bussiness

import "github.com/dragranzer/Golang-Mini-Project/features/detail_book"

type detail_bookUsecase struct {
	detail_bookData detail_book.Data
}

func NewDetailBookBussiness(detData detail_book.Data) detail_book.Bussiness {
	return &detail_bookUsecase{
		detail_bookData: detData,
	}
}

func (du *detail_bookUsecase) CreateData(data detail_book.Core) (resp detail_book.Core, err error) {

	data, err = du.detail_bookData.InsertData(data)
	if err != nil {
		return detail_book.Core{}, err
	}
	return data, nil
}
