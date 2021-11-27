package data

import (
	"github.com/dragranzer/Golang-Mini-Project/features/books"
	"gorm.io/gorm"
)

type mysqlBookRepository struct {
	Conn *gorm.DB
}

func NewBookRepository(conn *gorm.DB) books.Data {
	return &mysqlBookRepository{
		Conn: conn,
	}
}

func (ar *mysqlBookRepository) InsertData(data books.Core) (resp books.Core, err error) {
	// fmt.Println("Data ========", data)
	recordData := fromCore(data)
	// fmt.Println("Record data ======== ", recordData)
	err = ar.Conn.Create(&recordData).Error
	if err != nil {
		return books.Core{}, err
	}
	return data, nil
}

func (ar *mysqlBookRepository) SelectAllData() (resp []books.Core) {
	record := []Book{}

	if err := ar.Conn.Preload("Kategori").Find(&record).Error; err != nil {
		return []books.Core{}
	}

	// jangan lupa ditranlasiin ke core
	return toCoreList(record)
}
