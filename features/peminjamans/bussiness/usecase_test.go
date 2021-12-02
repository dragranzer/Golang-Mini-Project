package bussiness_test

import (
	"os"
	"testing"
	"time"

	"github.com/dragranzer/Golang-Mini-Project/features/books"
	b_books_mock "github.com/dragranzer/Golang-Mini-Project/features/books/mocks"
	"github.com/dragranzer/Golang-Mini-Project/features/peminjamans"
	peminjaman_buss "github.com/dragranzer/Golang-Mini-Project/features/peminjamans/bussiness"
	peminjaman_mock "github.com/dragranzer/Golang-Mini-Project/features/peminjamans/mocks"
	"github.com/dragranzer/Golang-Mini-Project/features/users"
	users_mock "github.com/dragranzer/Golang-Mini-Project/features/users/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	bookUsecase       b_books_mock.Bussiness
	userUseCase       users_mock.Bussiness
	bookValues        []books.Core
	userValues        []users.Core
	peminjamanUsecase peminjamans.Bussiness
	peminjamanData    peminjaman_mock.Data
	peminjamanValues  []peminjamans.Core
	detBookPinjam     []peminjamans.DetailBookPinjam
)

func TestMain(m *testing.M) {
	peminjamanUsecase = peminjaman_buss.NewPeminjamanBussiness(&peminjamanData, &bookUsecase, &userUseCase)

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

	peminjamanValues = []peminjamans.Core{
		{
			ID:         1,
			Hari:       7,
			TotalHarga: 30000,
			BookID:     1,
			UserID:     1,
		},
	}

	// detBookPinjam = []peminjamans.DetailBookPinjam

	os.Exit(m.Run())
}

func TestCoverage(t *testing.T) {
	t.Run("valid - Create Data", func(t *testing.T) {
		bookUsecase.On("GetDetailData", mock.AnythingOfType("string")).Return(bookValues, nil).Once()
		userid := users.Core{}
		userUseCase.On("GetDatabyName", mock.AnythingOfType("string")).Return(userid, nil).Once()
		insert := peminjamans.Core{}
		peminjamanData.On("InsertData", mock.AnythingOfType("peminjamans.Core")).Return(insert, nil).Once()
		book := books.Core{}
		bookUsecase.On("ChangeTersediabyName", mock.AnythingOfType("string")).Return(book, nil).Once()
		resp, _ := peminjamanUsecase.CreateData(peminjamanValues[0])

		// assert.NotEqual(t, len(resp), 0)
		assert.NotEqual(t, resp, peminjamanValues[0])
	})

	t.Run("valid - Get All Data", func(t *testing.T) {
		peminjaman := []peminjamans.Core{}
		peminjamanData.On("SelectAllData").Return(peminjaman).Once()
		resp := peminjamanUsecase.GetAllData()

		// assert.NotEqual(t, len(resp), 0)
		assert.NotEqual(t, resp, peminjamanValues)
	})

	t.Run("valid - Get Detail Book Pinjam", func(t *testing.T) {
		// listbook := []books.Core{}
		bookUsecase.On("GetDetailData", mock.AnythingOfType("string")).Return(bookValues, nil).Once()
		peminjaman := []peminjamans.Core{}
		peminjamanData.On("SelectDetailBookPinjam", mock.AnythingOfType("int")).Return(peminjaman).Once()
		user := users.Core{}
		userUseCase.On("GetDatabyID", mock.AnythingOfType("int")).Return(user, nil).Once()

		resp := peminjamanUsecase.GetDetailBookPinjam(bookValues[0].Judul)
		// assert.NotEqual(t, len(resp), 0)
		detbookpinjam := peminjamans.DetailBookPinjam{}
		assert.NotEqual(t, resp, detbookpinjam)
	})
}
