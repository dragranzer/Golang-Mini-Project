package bussiness_test

import (
	"os"
	"testing"
	"time"

	"github.com/dragranzer/Golang-Mini-Project/features/authors"
	b_authors_mock "github.com/dragranzer/Golang-Mini-Project/features/authors/mocks"
	"github.com/dragranzer/Golang-Mini-Project/features/books"
	b_books "github.com/dragranzer/Golang-Mini-Project/features/books/bussiness"
	b_books_mock "github.com/dragranzer/Golang-Mini-Project/features/books/mocks"
	"github.com/dragranzer/Golang-Mini-Project/features/detail_book"
	b_detBook_mock "github.com/dragranzer/Golang-Mini-Project/features/detail_book/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	bookData       b_books_mock.Data
	bookUsecase    books.Bussiness
	authorsUsecase b_authors_mock.Bussiness
	detBookUsecase b_detBook_mock.Bussiness
	Authors        []books.AuthorCore
	Authors2       []books.AuthorCore
	authorValues   []authors.Core

	bookValues  []books.Core
	bookValues2 []books.Core
)

func TestMain(m *testing.M) {

	bookUsecase = b_books.NewBookBussiness(&bookData, &authorsUsecase, &detBookUsecase)

	bookValues = []books.Core{
		{
			ID:          1,
			Judul:       "buku1",
			Tersedia:    100,
			PublishedAt: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
			Like:        29,
			Harga:       30000,
		},
	}

	bookValues2 = []books.Core{
		{
			ID:          1,
			Judul:       "buku1",
			Tersedia:    100,
			PublishedAt: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
			Like:        29,
			Harga:       30000,
			Authors: []books.AuthorCore{
				{Nama: "tooru"},
				{Nama: "elma"},
				{Nama: "Dazai"},
			},
		},
	}

	Authors = []books.AuthorCore{
		{Nama: "Nakahara"},
		{Nama: "Chuuya"},
	}

	Authors2 = []books.AuthorCore{
		{Nama: "tooru"},
		{Nama: "elma"},
		{Nama: "Dazai"},
	}

	authorValues = []authors.Core{
		{
			ID:   1,
			Nama: "Nakahara",
		},
	}
	os.Exit(m.Run())
}

func TestGetAllData(t *testing.T) {
	t.Run("valid - get all data", func(t *testing.T) {
		bookData.On("SelectAllData").Return(bookValues).Once()
		resp := bookUsecase.GetAllData()

		assert.NotEqual(t, len(resp), 0)
		assert.Equal(t, resp[0].ID, bookValues[0].ID)
	})

	t.Run("valid - insert data", func(t *testing.T) {
		bookData.On("InsertData", bookValues[0]).Return(bookValues[0], nil).Once()
		resp, _ := bookUsecase.CreateData(bookValues[0])

		// assert.NotEqual(t, len(resp), 0)
		assert.Equal(t, resp.ID, bookValues[0].ID)
	})

	t.Run("valid - get data by id", func(t *testing.T) {
		bookData.On("SelectbyID", bookValues[0].ID).Return(bookValues[0], nil).Once()
		// bookData.On("SelectData", bookValues[0].Judul).Return(bookValues, nil).Once()
		resp, _ := bookUsecase.GetDetailDatabyID(bookValues[0].ID)

		// assert.NotEqual(t, len(resp), 0)
		assert.Equal(t, resp, bookValues[0])
	})

	t.Run("valid - Update tersedia", func(t *testing.T) {
		bookData.On("UpdateTersediabyName", bookValues[0].Judul).Return(bookValues[0], nil).Once()
		resp, _ := bookUsecase.ChangeTersediabyName(bookValues[0].Judul)

		// assert.NotEqual(t, len(resp), 0)
		assert.Equal(t, resp, bookValues[0])
	})

	t.Run("valid - Get Detail Data", func(t *testing.T) {
		bookData.On("SelectData", bookValues[0].Judul).Return(bookValues, nil).Once()
		arr := []int{1, 2}

		detBookUsecase.On("GetAuthorbyBookID", bookValues[0].ID).Return(arr).Once()
		authorsUsecase.On("GetDetailDatabyID", mock.AnythingOfType("int")).Return(authorValues[0], nil).Times(len(arr))

		resp, _ := bookUsecase.GetDetailData(bookValues[0].Judul)
		assert.NotEqual(t, resp, bookValues[0])
	})

	t.Run("valid - insert data v2", func(t *testing.T) {
		listAuthor := []authors.Core{}
		authorsUsecase.On("GetDetailData", mock.AnythingOfType("string")).Return(listAuthor, nil).Times(len(bookValues2[0].Authors))
		newAuthor := authors.Core{}
		authorsUsecase.On("CreateData", mock.AnythingOfType("authors.Core")).Return(newAuthor, nil).Times(len(bookValues2[0].Authors))
		bookData.On("InsertData", bookValues2[0]).Return(bookValues2[0], nil).Times(len(bookValues[0].Authors))
		newDetBook := detail_book.Core{}
		detBookUsecase.On("CreateData", mock.AnythingOfType("detail_book.Core")).Return(newDetBook, nil).Times(len(bookValues2[0].Authors))

		resp, _ := bookUsecase.CreateData(bookValues2[0])
		assert.Equal(t, resp, bookValues2[0])
	})

}

// func TestCreateData(t *testing.T) {
// 	t.Run("valid - insert data", func(t *testing.T) {
// 		bookData.On("InsertData", bookValues[0]).Return(bookValues[0], nil).Once()
// 		resp, _ := bookUsecase.CreateData(bookValues[0])

// 		// assert.NotEqual(t, len(resp), 0)
// 		assert.Equal(t, resp.ID, bookValues[0].ID)
// 	})
// }

// func TestGetDetailDatabyID(t *testing.T) {
// 	t.Run("valid - get data by id", func(t *testing.T) {
// 		bookData.On("SelectbyID", bookValues[0].ID).Return(bookValues[0], nil).Once()
// 		resp, _ := bookUsecase.GetDetailDatabyID(bookValues[0].ID)

// 		// assert.NotEqual(t, len(resp), 0)
// 		assert.Equal(t, resp, bookValues[0])
// 	})
// }
// func TestGetDetailData(t *testing.T) {
// 	t.Run("valid - insert data v2", func(t *testing.T) {
// 		listAuthor := []authors.Core{}
// 		authorsUsecase.On("GetDetailData", mock.AnythingOfType("string")).Return(listAuthor, nil).Times(len(bookValues2[0].Authors))
// 		newAuthor := authors.Core{}
// 		authorsUsecase.On("CreateData", mock.AnythingOfType("authors.Core")).Return(newAuthor, nil).Times(len(bookValues[0].Authors))
// 		bookData.On("InsertData", bookValues2[0]).Return(bookValues2[0], nil).Times(len(bookValues[0].Authors))
// 		newDetBook := detail_book.Core{}
// 		detBookUsecase.On("CreateData", mock.AnythingOfType("detail_book.Core")).Return(newDetBook, nil).Times(len(bookValues[0].Authors))

// 		resp, _ := bookUsecase.CreateData(bookValues2[0])
// 		assert.Equal(t, resp, bookValues2[0])
// 	})
// }
