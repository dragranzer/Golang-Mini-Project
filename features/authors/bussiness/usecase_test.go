package bussiness_test

import (
	"os"
	"testing"

	"github.com/dragranzer/Golang-Mini-Project/features/authors"
	author_buss "github.com/dragranzer/Golang-Mini-Project/features/authors/bussiness"
	author_mock "github.com/dragranzer/Golang-Mini-Project/features/authors/mocks"
	"github.com/stretchr/testify/assert"
)

var (
	authorUsecase authors.Bussiness
	authorData    author_mock.Data
	authorValues  []authors.Core
)

func TestMain(m *testing.M) {
	authorUsecase = author_buss.NewAuthorBussiness(&authorData)

	authorValues = []authors.Core{
		{
			ID:   1,
			Nama: "Nakahara",
		},
	}

	os.Exit(m.Run())
}

func TestCoverage(t *testing.T) {
	t.Run("valid - Create Data", func(t *testing.T) {
		author := authors.Core{}
		authorData.On("InsertData", authorValues[0]).Return(author, nil).Once()
		resp, _ := authorUsecase.CreateData(authorValues[0])

		// assert.NotEqual(t, len(resp), 0)
		assert.NotEqual(t, resp, authorValues[0])
	})

	t.Run("valid - Get All Data", func(t *testing.T) {
		authorData.On("SelectAllData").Return(authorValues).Once()
		resp := authorUsecase.GetAllData()

		// assert.NotEqual(t, len(resp), 0)
		assert.Equal(t, resp, authorValues)
	})

	t.Run("valid - Get Data by Name", func(t *testing.T) {
		authorData.On("SelectData", authorValues[0].Nama).Return(authorValues, nil).Once()
		resp, _ := authorUsecase.GetDetailData(authorValues[0].Nama)

		// assert.NotEqual(t, len(resp), 0)
		assert.Equal(t, resp, authorValues)
	})

	t.Run("valid - Get Detail Data by ID(", func(t *testing.T) {
		authorData.On("SelectDatabyID", authorValues[0].ID).Return(authorValues[0], nil).Once()
		resp, _ := authorUsecase.GetDetailDatabyID((authorValues[0].ID))

		// assert.NotEqual(t, len(resp), 0)
		assert.NotEqual(t, resp, authorValues)
	})
}
