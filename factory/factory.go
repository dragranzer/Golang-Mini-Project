package factory

import (
	"github.com/dragranzer/Golang-Mini-Project/config"
	_book_bussiness "github.com/dragranzer/Golang-Mini-Project/features/books/bussiness"
	_book_data "github.com/dragranzer/Golang-Mini-Project/features/books/data"
	_book_presentation "github.com/dragranzer/Golang-Mini-Project/features/books/presentation"

	_author_bussiness "github.com/dragranzer/Golang-Mini-Project/features/authors/bussiness"
	_author_data "github.com/dragranzer/Golang-Mini-Project/features/authors/data"
	_author_presentation "github.com/dragranzer/Golang-Mini-Project/features/authors/presentation"

	_user_bussiness "github.com/dragranzer/Golang-Mini-Project/features/users/bussiness"
	_user_data "github.com/dragranzer/Golang-Mini-Project/features/users/data"
	_user_presentation "github.com/dragranzer/Golang-Mini-Project/features/users/presentation"

	_detail_book_bussiness "github.com/dragranzer/Golang-Mini-Project/features/detail_book/bussiness"
	_detail_book_data "github.com/dragranzer/Golang-Mini-Project/features/detail_book/data"
	_detail_book_presentation "github.com/dragranzer/Golang-Mini-Project/features/detail_book/presentation"

	_peminjaman_bussiness "github.com/dragranzer/Golang-Mini-Project/features/peminjamans/bussiness"
	_peminjaman_data "github.com/dragranzer/Golang-Mini-Project/features/peminjamans/data"
	_peminjaman_presentation "github.com/dragranzer/Golang-Mini-Project/features/peminjamans/presentation"

	_favorite_bussiness "github.com/dragranzer/Golang-Mini-Project/features/favorites/bussiness"
	_favorite_data "github.com/dragranzer/Golang-Mini-Project/features/favorites/data"
	_favorite_presentation "github.com/dragranzer/Golang-Mini-Project/features/favorites/presentation"

	_review_bussiness "github.com/dragranzer/Golang-Mini-Project/features/reviews/bussiness"
	_review_data "github.com/dragranzer/Golang-Mini-Project/features/reviews/data"
	_review_presentation "github.com/dragranzer/Golang-Mini-Project/features/reviews/presentation"
)

type Presenter struct {
	BookPresentation       *_book_presentation.BooksHandler
	AuthorPresentation     *_author_presentation.AuthorsHandler
	UserPresentation       *_user_presentation.UsersHandler
	DetailBookPresentation *_detail_book_presentation.Detail_bookHandler
	PeminjamanPresentation *_peminjaman_presentation.PeminjamansHandler
	FavoritePresentation   *_favorite_presentation.FavoritesHandler
	ReviewPresentation     *_review_presentation.ReviewsHandler
}

func Init() Presenter {

	bookData := _book_data.NewBookRepository(config.DB)
	authorData := _author_data.NewAuthorRepository(config.DB)
	userData := _user_data.NewUserRepository(config.DB)
	detail_bookData := _detail_book_data.NewDetailBookRepository(config.DB)
	peminjamanData := _peminjaman_data.NewPeminjamanRepository(config.DB)
	favoriteData := _favorite_data.NewFavoriteRepository(config.DB)
	reviewData := _review_data.NewReviewRepository(config.DB)

	detail_bookBussiness := _detail_book_bussiness.NewDetailBookBussiness(detail_bookData)
	authorBussiness := _author_bussiness.NewAuthorBussiness(authorData)
	bookBussiness := _book_bussiness.NewBookBussiness(bookData, authorBussiness, detail_bookBussiness)
	userBussiness := _user_bussiness.NewUserBussiness(userData, bookBussiness)
	peminjamanBussiness := _peminjaman_bussiness.NewPeminjamanBussiness(peminjamanData, bookBussiness, userBussiness)
	favoriteBussiness := _favorite_bussiness.NewFavoriteBussiness(favoriteData, bookBussiness, userBussiness)
	reviewBussiness := _review_bussiness.NewReviewBussiness(reviewData, bookBussiness, userBussiness)

	// authorData := _author_data.NewAuthorRepository(config.DB)
	// authorBussiness := _author_bussiness.NewAuthorBussiness(authorData)

	return Presenter{
		BookPresentation:       _book_presentation.NewBookHandler(bookBussiness),
		AuthorPresentation:     _author_presentation.NewAuthorHandler(authorBussiness),
		UserPresentation:       _user_presentation.NewUserHandler(userBussiness),
		DetailBookPresentation: _detail_book_presentation.Newdetail_bookHandler(detail_bookBussiness),
		PeminjamanPresentation: _peminjaman_presentation.NewPeminjamanHandler(peminjamanBussiness),
		FavoritePresentation:   _favorite_presentation.NewFavoriteHandler(favoriteBussiness),
		ReviewPresentation:     _review_presentation.NewReviewHandler(reviewBussiness),
	}
}
