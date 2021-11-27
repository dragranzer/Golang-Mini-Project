package bussiness

import (
	"github.com/dragranzer/Golang-Mini-Project/features/books"
)

type booksUsecase struct {
	bookData books.Data
}

func NewBookBussiness(artData books.Data) books.Bussiness {
	return &booksUsecase{
		bookData: artData,
	}
}

func (bu *booksUsecase) CreateData(data books.Core) (resp books.Core, err error) {
	// if err := bu.validate.Struct(data); err != nil {
	// 	return books.Core{}, err
	// }

	data, err = bu.bookData.InsertData(data)
	if err != nil {
		return books.Core{}, err
	}
	return data, nil
}

func (bu *booksUsecase) GetAllData() (resp []books.Core) {
	resp = bu.bookData.SelectAllData()
	return
}
