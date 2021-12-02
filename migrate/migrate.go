package migrate

import (
	"time"

	"github.com/dragranzer/Golang-Mini-Project/config"
	_author_data "github.com/dragranzer/Golang-Mini-Project/features/authors/data"
	_book_data "github.com/dragranzer/Golang-Mini-Project/features/books/data"
	_favorite_data "github.com/dragranzer/Golang-Mini-Project/features/favorites/data"
	_peminjaman_data "github.com/dragranzer/Golang-Mini-Project/features/peminjamans/data"
	_review_data "github.com/dragranzer/Golang-Mini-Project/features/reviews/data"
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
	if err := config.DB.Exec("DROP TABLE IF EXISTS reviews").Error; err != nil {
		panic(err)
	}

	if err := config.DB.Exec("DROP TABLE IF EXISTS favorites").Error; err != nil {
		panic(err)
	}

	if err := config.DB.Exec("DROP TABLE IF EXISTS kategoris").Error; err != nil {
		panic(err)
	}

	if err := config.DB.Exec("DROP TABLE IF EXISTS detail_books").Error; err != nil {
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

	if err := config.DB.Exec("DROP TABLE IF EXISTS peminjamen").Error; err != nil {
		panic(err)
	}

	config.DB.AutoMigrate(&_book_data.Book{}, &_book_data.Kategori{},
		&_author_data.Author{}, &_book_data.Author{}, &_user_data.User{},
		&_peminjaman_data.Peminjaman{}, &_favorite_data.Favorite{}, &_review_data.Review{})

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

	// author1 := _author_data.Author{
	// 	Nama: "Dazai",
	// 	Books: []_author_data.Book{
	// 		{Judul: "buku2"},
	// 	},
	// }

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

	peminjaman1 := _peminjaman_data.Peminjaman{
		Hari:       7,
		TotalHarga: 30000,
		BookID:     1,
		UserID:     1,
	}

	favorite1 := _favorite_data.Favorite{
		BookID: 1,
		UserID: 2,
	}

	favorite2 := _favorite_data.Favorite{
		BookID: 2,
		UserID: 2,
	}

	review1 := _review_data.Review{
		BookID: 1,
		UserID: 2,
		Review: "Keren banget buku nya, insightfull",
	}

	review2 := _review_data.Review{
		BookID: 1,
		UserID: 1,
		Review: "keren gan",
	}

	review3 := _review_data.Review{
		BookID: 2,
		UserID: 1,
		Review: "keren gan, sangar",
	}

	review4 := _review_data.Review{
		BookID: 2,
		UserID: 2,
		Review: "sama kaya user1",
	}

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
	if err := config.DB.Create(&user1).Error; err != nil {
		panic(err)
	}
	if err := config.DB.Create(&user2).Error; err != nil {
		panic(err)
	}
	if err := config.DB.Create(&peminjaman1).Error; err != nil {
		panic(err)
	}
	if err := config.DB.Create(&favorite1).Error; err != nil {
		panic(err)
	}
	if err := config.DB.Create(&favorite2).Error; err != nil {
		panic(err)
	}
	if err := config.DB.Create(&review1).Error; err != nil {
		panic(err)
	}
	if err := config.DB.Create(&review2).Error; err != nil {
		panic(err)
	}
	if err := config.DB.Create(&review3).Error; err != nil {
		panic(err)
	}
	if err := config.DB.Create(&review4).Error; err != nil {
		panic(err)
	}
}
