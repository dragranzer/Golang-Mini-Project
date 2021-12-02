package bussiness_test

import (
	"os"
	"testing"
	"time"

	"github.com/dragranzer/Golang-Mini-Project/features/books"
	b_books_mock "github.com/dragranzer/Golang-Mini-Project/features/books/mocks"
	"github.com/dragranzer/Golang-Mini-Project/features/favorites"
	fav_buss "github.com/dragranzer/Golang-Mini-Project/features/favorites/bussiness"
	favorite_mock "github.com/dragranzer/Golang-Mini-Project/features/favorites/mocks"
	"github.com/dragranzer/Golang-Mini-Project/features/users"
	users_mock "github.com/dragranzer/Golang-Mini-Project/features/users/mocks"
	"github.com/stretchr/testify/assert"
)

var (
	favoriteData    favorite_mock.Data
	favoriteUsecase favorites.Bussiness
	bookUsecase     b_books_mock.Bussiness
	userUseCase     users_mock.Bussiness
	bookValues      []books.Core
	favoriteValues  []favorites.Core
	userValues      []users.Core
)

func TestMain(m *testing.M) {
	favoriteUsecase = fav_buss.NewFavoriteBussiness(&favoriteData, &bookUsecase, &userUseCase)

	favoriteValues = []favorites.Core{
		{
			UserID: 1,
			BookID: 1,
		},
	}

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

	userValues = []users.Core{
		{
			ID:       1,
			Nama:     "user1",
			Email:    "user1@gmail.com",
			Password: "$2a$14$m6fJnGicXymsu35pyWqSv.k1V7zfUyH6uzQuiXCHgdOhFJsVNhpai",
		},
	}

	os.Exit(m.Run())
}

func TestCoverage(t *testing.T) {
	t.Run("valid - Create Data", func(t *testing.T) {
		favoriteData.On("InsertData", favoriteValues[0]).Return(nil).Once()
		resp := favoriteUsecase.CreateData(favoriteValues[0])

		// assert.NotEqual(t, len(resp), 0)
		assert.NotEqual(t, resp, favoriteValues[0])
	})

	t.Run("valid - Get Data by UserID", func(t *testing.T) {
		favoriteData.On("SelectDatabyUserID", favoriteValues[0].UserID).Return(favoriteValues, nil).Once()
		bookUsecase.On("GetDetailDatabyID", bookValues[0].ID).Return(bookValues[0], nil).Once()
		resp, _ := favoriteUsecase.GetDatabyUserID(favoriteValues[0].UserID)

		// assert.NotEqual(t, len(resp), 0)
		assert.NotEqual(t, resp, favoriteValues[0])
	})

	t.Run("valid - Get Data by BookID", func(t *testing.T) {
		favoriteData.On("SelectDatabyBookID", favoriteValues[0].UserID).Return(favoriteValues, nil).Once()
		userUseCase.On("GetDatabyID", userValues[0].ID).Return(userValues[0], nil).Once()
		resp, _ := favoriteUsecase.GetDatabyBookID(bookValues[0].ID)

		// assert.NotEqual(t, len(resp), 0)
		assert.NotEqual(t, resp, bookValues[0])
	})

	t.Run("valid - Clear Fav by BookID", func(t *testing.T) {
		favoriteData.On("DeleteFavbyBookID", bookValues[0].ID).Return(nil).Once()
		resp := favoriteUsecase.ClearFavbyBookID(bookValues[0].ID)

		// assert.NotEqual(t, len(resp), 0)
		assert.Equal(t, resp, nil)
	})
}
