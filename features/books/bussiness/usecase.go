package bussiness

import (
	"fmt"

	"github.com/dragranzer/Golang-Mini-Project/features/authors"
	"github.com/dragranzer/Golang-Mini-Project/features/books"
	"github.com/dragranzer/Golang-Mini-Project/features/detail_book"
)

type booksUsecase struct {
	bookData         books.Data
	authorBussiness  authors.Bussiness
	detBookBussiness detail_book.Bussiness
}

func NewBookBussiness(bookData books.Data, ab authors.Bussiness,
	dBus detail_book.Bussiness) books.Bussiness {
	return &booksUsecase{
		bookData:         bookData,
		authorBussiness:  ab,
		detBookBussiness: dBus,
	}
}

func (bu *booksUsecase) CreateData(data books.Core) (resp books.Core, err error) {
	// if err := bu.validate.Struct(data); err != nil {
	// 	return books.Core{}, err
	// }

	ListIDAuthor := []int{}

	for _, value := range data.Authors {

		author, _ := bu.authorBussiness.GetDetailData(value.Nama)
		if len(author) == 0 {
			newAuthor := authors.Core{
				Nama: value.Nama,
			}
			newAuthor, _ = bu.authorBussiness.CreateData(newAuthor)
			fmt.Println(newAuthor)
			ListIDAuthor = append(ListIDAuthor, newAuthor.ID)
		} else {
			for _, authorCore := range author {
				fmt.Println("list of auhor", authorCore.ID)
				ListIDAuthor = append(ListIDAuthor, authorCore.ID)
			}
		}
	}
	// fmt.Println("list of author", author)

	data, err = bu.bookData.InsertData(data)
	fmt.Println("bussiness book", data)
	if err != nil {
		return books.Core{}, err
	}

	// for _, authorid := range ListIDAuthor{

	// }
	for _, authorid := range ListIDAuthor {
		detBook := detail_book.Core{
			BookID:   data.ID,
			AuthorID: authorid,
		}
		bu.detBookBussiness.CreateData(detBook)
	}
	return data, nil
}

func (bu *booksUsecase) GetAllData() (resp []books.Core) {
	resp = bu.bookData.SelectAllData()
	return resp
}

func (bu *booksUsecase) GetDetailData(judul string) (resp []books.Core, err error) {
	resp, err = bu.bookData.SelectData(judul)
	listAuthorID := bu.detBookBussiness.GetAuthorbyBookID(resp[0].ID)
	// listPeminjam :=
	// fmt.Println("lsit author id", listAuthorID)
	listAuthorCore := []authors.Core{}
	for _, value := range listAuthorID {
		authorData, _ := bu.authorBussiness.GetDetailDatabyID(value)
		listAuthorCore = append(listAuthorCore, authorData)
	}
	fmt.Println(listAuthorCore)
	listAuthor := []books.AuthorCore{}

	for _, value := range listAuthorCore {
		author := books.AuthorCore{
			ID:   value.ID,
			Nama: value.Nama,
		}
		listAuthor = append(listAuthor, author)
	}
	resp[0] = books.Core{
		ID:          resp[0].ID,
		Judul:       resp[0].Judul,
		PublishedAt: resp[0].PublishedAt,
		Harga:       resp[0].Harga,
		Authors:     listAuthor,
	}
	if err != nil {
		return []books.Core{}, err
	}

	return resp, nil
}

func (bu *booksUsecase) ChangeTersediabyName(name string) (resp books.Core, err error) {

	fmt.Println("name masuk ", name)
	resp, err = bu.bookData.UpdateTersediabyName(name)
	if err != nil {
		return books.Core{}, err
	}

	return
}

func (bu *booksUsecase) GetDetailDatabyID(id int) (resp books.Core, err error) {
	resp, err = bu.bookData.SelectbyID(id)
	if err != nil {
		return books.Core{}, err
	}
	return
}

func (bu *booksUsecase) ClearDatabyBookID(id int) (err error) {

	err = bu.bookData.DeleteDatabyBookID(id)
	return
}
