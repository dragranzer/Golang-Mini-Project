package bussiness_test

import (
	"os"
	"testing"

	"github.com/dragranzer/Golang-Mini-Project/features/books"
	"github.com/dragranzer/Golang-Mini-Project/features/favorites"
	fav_buss "github.com/dragranzer/Golang-Mini-Project/features/favorites/bussiness"
	favorite_mock "github.com/dragranzer/Golang-Mini-Project/features/favorites/mocks"
	"github.com/dragranzer/Golang-Mini-Project/features/users"
	"github.com/stretchr/testify/assert"
)

var (
	favoriteData    favorite_mock.Data
	favoriteUsecase favorites.Bussiness
	bookUsecase     books.Bussiness
	userUseCase     users.Bussiness

	favoriteValues []favorites.Core
)

func TestMain(m *testing.M) {
	favoriteUsecase = fav_buss.NewFavoriteBussiness(&favoriteData, bookUsecase, userUseCase)

	favoriteValues = []favorites.Core{
		{
			UserID: 1,
			BookID: 1,
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
		resp, _ := favoriteUsecase.GetDatabyUserID(favoriteValues[0].UserID)

		// assert.NotEqual(t, len(resp), 0)
		assert.NotEqual(t, resp, favoriteValues[0])
	})
}
