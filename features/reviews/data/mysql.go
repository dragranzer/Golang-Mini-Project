package data

import (
	"github.com/dragranzer/Golang-Mini-Project/features/reviews"
	"gorm.io/gorm"
)

type mysqlReviewRepository struct {
	Conn *gorm.DB
}

func NewReviewRepository(conn *gorm.DB) reviews.Data {
	return &mysqlReviewRepository{
		Conn: conn,
	}
}

func (ar *mysqlReviewRepository) InsertData(data reviews.Core) (err error) {
	record := fromCore(data)
	err = ar.Conn.Create(&record).Error
	return
}

func (ar *mysqlReviewRepository) SelectDatabyBookID(id int) (resp []reviews.Core, err error) {

	record := []Review{}
	if err = ar.Conn.Where("book_id = ?", id).Find(&record).Error; err != nil {
		return []reviews.Core{}, err
	}
	return toCoreList(record), nil
}
