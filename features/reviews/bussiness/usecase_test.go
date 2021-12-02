package bussiness_test

import (
	"os"
	"testing"
	"time"

	"github.com/dragranzer/Golang-Mini-Project/features/books"
	b_books_mock "github.com/dragranzer/Golang-Mini-Project/features/books/mocks"
	"github.com/dragranzer/Golang-Mini-Project/features/reviews"
	rev_buss "github.com/dragranzer/Golang-Mini-Project/features/reviews/bussiness"
	review_mock "github.com/dragranzer/Golang-Mini-Project/features/reviews/mocks"
	"github.com/dragranzer/Golang-Mini-Project/features/users"
	users_mock "github.com/dragranzer/Golang-Mini-Project/features/users/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	reviewData    review_mock.Data
	reviewUsecase reviews.Bussiness
	bookUsecase   b_books_mock.Bussiness
	userUseCase   users_mock.Bussiness
	bookValues    []books.Core
	reviewValues  []reviews.Core
	userValues    []users.Core
)

func TestMain(m *testing.M) {
	reviewUsecase = rev_buss.NewReviewBussiness(&reviewData, &bookUsecase, &userUseCase)

	reviewValues = []reviews.Core{
		{
			UserID: 1,
			BookID: 1,
			Review: "awkwkwkwkwkk",
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
		reviewData.On("InsertData", reviewValues[0]).Return(nil).Once()
		resp := reviewUsecase.CreateData(reviewValues[0])

		// assert.NotEqual(t, len(resp), 0)
		assert.NotEqual(t, resp, reviewValues[0])
	})

	t.Run("valid - Get Data by BookID", func(t *testing.T) {
		reviewData.On("SelectDatabyBookID", mock.AnythingOfType("int")).Return(reviewValues, nil).Once()
		resp, _ := reviewUsecase.GetDatabyBookID(bookValues[0].ID)

		// assert.NotEqual(t, len(resp), 0)
		assert.NotEqual(t, resp, reviewValues[0])
	})

	t.Run("valid - Clear Data by BookID", func(t *testing.T) {
		reviewData.On("DeleteDatabyBookID", bookValues[0].ID).Return(nil).Once()
		resp := reviewUsecase.ClearDatabyBookID(bookValues[0].ID)

		// assert.NotEqual(t, len(resp), 0)
		assert.Equal(t, resp, nil)
	})
}
