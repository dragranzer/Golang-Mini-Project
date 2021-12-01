package bussiness_test

import (
	"os"
	"testing"

	"github.com/dragranzer/Golang-Mini-Project/features/detail_book"
	b_det_book "github.com/dragranzer/Golang-Mini-Project/features/detail_book/bussiness"
	b_det_book_mock "github.com/dragranzer/Golang-Mini-Project/features/detail_book/mocks"
	"github.com/stretchr/testify/assert"
)

var (
	detail_bookData    b_det_book_mock.Data
	detail_bookUsecase detail_book.Bussiness

	detail_bookValues []detail_book.Core
)

func TestMain(m *testing.M) {
	detail_bookUsecase = b_det_book.NewDetailBookBussiness(&detail_bookData)

	detail_bookValues = []detail_book.Core{
		{
			AuthorID: 1,
			BookID:   1,
		},
	}

	os.Exit(m.Run())
}

func TestCoverage(t *testing.T) {
	t.Run("valid - Create Data", func(t *testing.T) {
		detail_bookData.On("InsertData", detail_bookValues[0]).Return(detail_bookValues[0], nil).Once()
		resp, _ := detail_bookUsecase.CreateData(detail_bookValues[0])

		// assert.NotEqual(t, len(resp), 0)
		assert.Equal(t, resp, detail_bookValues[0])
	})

	t.Run("valid - Get Author by BookID", func(t *testing.T) {
		detail_bookData.On("SelectAuthorbyBookID", detail_bookValues[0].BookID).Return([]int{1, 2}, nil).Once()
		resp := detail_bookUsecase.GetAuthorbyBookID(detail_bookValues[0].BookID)

		// assert.NotEqual(t, len(resp), 0)
		assert.Equal(t, resp, []int{1, 2})
	})
}
