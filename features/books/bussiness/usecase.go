package bussiness

import (
	"fmt"

	"github.com/dragranzer/Golang-Mini-Project/features/books"
)

type booksUsecase struct {
	bookData books.Data
}

func NewAricleBussiness(artData books.Data) books.Bussiness {
	fmt.Println("Usecase Bussiness")
	fmt.Println(artData)
	return &booksUsecase{
		bookData: artData,
	}
}

func (bu *booksUsecase) CreateData(data books.Core) (resp books.Core, err error) {
	// if err := bu.validate.Struct(data); err != nil {
	// 	return books.Core{}, err
	// }
	fmt.Println("bussiness", data)
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
