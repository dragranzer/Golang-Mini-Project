package factory

import (
	"github.com/dragranzer/Golang-Mini-Project/config"

	_author_bussiness "github.com/dragranzer/Golang-Mini-Project/features/authors/bussiness"
	_author_data "github.com/dragranzer/Golang-Mini-Project/features/authors/data"
	_author_presentation "github.com/dragranzer/Golang-Mini-Project/features/authors/presentation"
)

type PresenterAuthor struct {
	AuthorPresentation _author_presentation.AuthorsHandler
}

func InitAuthor() PresenterAuthor {

	authorData := _author_data.NewAuthorRepository(config.DB)
	authorBussiness := _author_bussiness.NewAuthorBussiness(authorData)

	return PresenterAuthor{
		AuthorPresentation: *_author_presentation.NewAuthorHandler(authorBussiness),
	}
}
