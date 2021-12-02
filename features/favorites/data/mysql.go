package data

import (
	"github.com/dragranzer/Golang-Mini-Project/features/favorites"
	"gorm.io/gorm"
)

type mysqlFavoriteRepository struct {
	Conn *gorm.DB
}

func NewFavoriteRepository(conn *gorm.DB) favorites.Data {
	return &mysqlFavoriteRepository{
		Conn: conn,
	}
}

func (ar *mysqlFavoriteRepository) InsertData(data favorites.Core) (err error) {
	record := fromCore(data)
	err = ar.Conn.Create(&record).Error
	return
}

func (ar *mysqlFavoriteRepository) SelectDatabyUserID(id int) (resp []favorites.Core, err error) {

	record := []Favorite{}
	if err = ar.Conn.Where("user_id = ?", id).Find(&record).Error; err != nil {
		return []favorites.Core{}, err
	}
	return toCoreList(record), nil
}

func (ar *mysqlFavoriteRepository) SelectDatabyBookID(id int) (resp []favorites.Core, err error) {

	record := []Favorite{}
	if err = ar.Conn.Where("book_id = ?", id).Find(&record).Error; err != nil {
		return []favorites.Core{}, err
	}
	return toCoreList(record), nil
}

func (ar *mysqlFavoriteRepository) DeleteFavbyBookID(id int) (err error) {
	record := Favorite{}
	err = ar.Conn.Where("book_id = ?", id).Delete(&record).Error
	return
}
