package migrate

import (
	"time"

	"github.com/dragranzer/Golang-Mini-Project/config"
	_author_data "github.com/dragranzer/Golang-Mini-Project/features/authors/data"
	_book_data "github.com/dragranzer/Golang-Mini-Project/features/books/data"
)

func AutoMigrate() {
	if err := config.DB.Exec("DROP TABLE IF EXISTS kategoris").Error; err != nil {
		panic(err)
	}

	if err := config.DB.Exec("DROP TABLE IF EXISTS detail_book").Error; err != nil {
		panic(err)
	}

	if err := config.DB.Exec("DROP TABLE IF EXISTS books").Error; err != nil {
		panic(err)
	}

	if err := config.DB.Exec("DROP TABLE IF EXISTS authors").Error; err != nil {
		panic(err)
	}

	// if err := config.DB.Exec("DROP TABLE IF EXISTS detail_book").Error; err != nil {
	// 	panic(err)
	// }

	config.DB.AutoMigrate(&_book_data.Book{}, &_book_data.Kategori{},
		&_author_data.Author{}, &_author_data.Author{}, &_book_data.Author{})

	book1 := _book_data.Book{
		Judul:       "buku1",
		Tersedia:    100,
		PublishedAt: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
		Like:        29,
		Harga:       30000,
		Kategori: []_book_data.Kategori{
			{Nama: "sci-fi"},
			{Nama: "history"},
		},
		Author: []_book_data.Author{
			{Nama: "Nakahara"},
			{Nama: "Chuuya"},
		},
	}

	book2 := _book_data.Book{
		Judul:       "buku2",
		Tersedia:    100,
		PublishedAt: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
		Like:        29,
		Harga:       30000,
		Kategori: []_book_data.Kategori{
			{Nama: "sci-fi"},
			{Nama: "history"},
		},
		Author: []_book_data.Author{
			{Nama: "Osamu"},
			{Nama: "Dazai"},
		},
	}

	// author1 := _author_data.Author{
	// 	Nama: "Dazai",
	// 	Book: []_author_data.Book{
	// 		{Judul: "buku2"},
	// 	},
	// }

	// setelah dibuat, insert
	if err := config.DB.Create(&book1).Error; err != nil {
		panic(err)
	}
	if err := config.DB.Create(&book2).Error; err != nil {
		panic(err)
	}
	// if err := config.DB.Create(&author1).Error; err != nil {
	// 	panic(err)
	// }

}
