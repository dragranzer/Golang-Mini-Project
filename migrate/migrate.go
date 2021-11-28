package migrate

import (
	"time"

	"github.com/dragranzer/Golang-Mini-Project/config"
	_author_data "github.com/dragranzer/Golang-Mini-Project/features/authors/data"
	_book_data "github.com/dragranzer/Golang-Mini-Project/features/books/data"
	_user_data "github.com/dragranzer/Golang-Mini-Project/features/users/data"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

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

	if err := config.DB.Exec("DROP TABLE IF EXISTS users").Error; err != nil {
		panic(err)
	}

	// if err := config.DB.Exec("DROP TABLE IF EXISTS detail_book").Error; err != nil {
	// 	panic(err)
	// }

	config.DB.AutoMigrate(&_book_data.Book{}, &_book_data.Kategori{},
		&_author_data.Author{}, &_book_data.Author{}, &_user_data.User{})

	book1 := _book_data.Book{
		Judul:       "buku1",
		Tersedia:    100,
		PublishedAt: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
		Like:        29,
		Harga:       30000,
		Kategoris: []_book_data.Kategori{
			{Nama: "sci-fi"},
			{Nama: "history"},
		},
		Authors: []_book_data.Author{
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
		Kategoris: []_book_data.Kategori{
			{Nama: "sci-fi"},
			{Nama: "history"},
		},
		Authors: []_book_data.Author{
			{Nama: "Osamu"},
			{Nama: "Dazai"},
		},
	}

	author1 := _author_data.Author{
		Nama: "Dazai",
		Book: []_author_data.Book{
			{Judul: "buku2"},
		},
	}

	pass, _ := HashPassword("testing")
	user1 := _user_data.User{
		Nama:     "user1",
		Email:    "user1@gmail.com",
		Password: pass,
	}

	pass1, _ := HashPassword("pass12344")
	user2 := _user_data.User{
		Nama:     "user2",
		Email:    "user2@gmail.comasjf",
		Password: pass1,
	}

	// setelah dibuat, insert
	if err := config.DB.Create(&book1).Error; err != nil {
		panic(err)
	}
	if err := config.DB.Create(&book2).Error; err != nil {
		panic(err)
	}
	if err := config.DB.Create(&author1).Error; err != nil {
		panic(err)
	}
	if err := config.DB.Create(&user1).Error; err != nil {
		panic(err)
	}
	if err := config.DB.Create(&user2).Error; err != nil {
		panic(err)
	}

}
