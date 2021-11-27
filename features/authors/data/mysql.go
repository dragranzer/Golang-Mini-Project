package data

import (
	"github.com/dragranzer/Golang-Mini-Project/features/authors"
	"gorm.io/gorm"
)

type mysqlAuthorRepository struct {
	Conn *gorm.DB
}

func NewAuthorRepository(conn *gorm.DB) authors.Data {
	return &mysqlAuthorRepository{
		Conn: conn,
	}
}

func (ar *mysqlAuthorRepository) InsertData(data authors.Core) (resp authors.Core, err error) {
	// fmt.Println("Data ========", data)
	recordData := fromCore(data)
	// fmt.Println("Record data ======== ", recordData)
	err = ar.Conn.Create(&recordData).Error
	if err != nil {
		return authors.Core{}, err
	}
	return data, nil
}

func (ar *mysqlAuthorRepository) SelectAllData() (resp []authors.Core) {
	record := []Author{}

	if err := ar.Conn.Find(&record).Error; err != nil {
		return []authors.Core{}
	}

	// jangan lupa ditranlasiin ke core
	return toCoreList(record)
}
