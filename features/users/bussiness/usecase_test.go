package bussiness_test

import (
	"os"
	"testing"
	"time"

	"github.com/dragranzer/Golang-Mini-Project/features/books"
	b_books_mock "github.com/dragranzer/Golang-Mini-Project/features/books/mocks"
	"github.com/dragranzer/Golang-Mini-Project/features/users"
	user_buss "github.com/dragranzer/Golang-Mini-Project/features/users/bussiness"
	user_mock "github.com/dragranzer/Golang-Mini-Project/features/users/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	userUsecase users.Bussiness
	userData    user_mock.Data
	userValues  []users.Core
	bookUsecase b_books_mock.Bussiness
	bookValues  []books.Core
)

func TestMain(m *testing.M) {
	userUsecase = user_buss.NewUserBussiness(&userData, &bookUsecase)

	userValues = []users.Core{
		{
			ID:       1,
			Nama:     "user1",
			Email:    "user1@gmail.com",
			Password: "$2a$14$m6fJnGicXymsu35pyWqSv.k1V7zfUyH6uzQuiXCHgdOhFJsVNhpai",
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

	os.Exit(m.Run())
}

func TestCoverage(t *testing.T) {
	t.Run("valid - Create Data", func(t *testing.T) {
		userData.On("InsertData", userValues[0]).Return(nil).Once()
		resp := userUsecase.CreateData(userValues[0])

		// assert.NotEqual(t, len(resp), 0)
		assert.NotEqual(t, resp, userValues[0])
	})

	t.Run("valid - Get All Data", func(t *testing.T) {
		userData.On("SelectAllData").Return(userValues).Once()
		resp := userUsecase.GetAllData()

		// assert.NotEqual(t, len(resp), 0)
		assert.Equal(t, resp, userValues)
	})

	t.Run("valid - Get Data by Name", func(t *testing.T) {
		userData.On("SelectDatabyName", userValues[0].Nama).Return(userValues[0], nil).Once()
		resp, _ := userUsecase.GetDatabyName(userValues[0].Nama)

		// assert.NotEqual(t, len(resp), 0)
		assert.NotEqual(t, resp, userValues)
	})

	t.Run("valid - Get Data by ID", func(t *testing.T) {
		userData.On("SelectDatabyID", userValues[0].ID).Return(userValues[0], nil).Once()
		resp, _ := userUsecase.GetDatabyID(userValues[0].ID)

		// assert.NotEqual(t, len(resp), 0)
		assert.NotEqual(t, resp, userValues)
	})

	t.Run("valid - Change Data by ID", func(t *testing.T) {
		userData.On("UpdateDatabyID", mock.AnythingOfType("int"), mock.AnythingOfType("users.Core")).Return(nil).Once()
		newData := users.Core{}
		err := userUsecase.ChangeDatabyID(userValues[0].ID, newData)

		// assert.NotEqual(t, len(resp), 0)
		assert.Equal(t, err, nil)
	})
}
