package data

import (
	"fmt"

	"github.com/dragranzer/Golang-Mini-Project/features/detail_book"
	"gorm.io/gorm"
)

type mysqlDetailBookRepository struct {
	Conn *gorm.DB
}

func NewDetailBookRepository(conn *gorm.DB) detail_book.Data {
	return &mysqlDetailBookRepository{
		Conn: conn,
	}
}

func (ar *mysqlDetailBookRepository) InsertData(data detail_book.Core) (resp detail_book.Core, err error) {
	// fmt.Println("Data ========", data)
	recordData := fromCore(data)
	// fmt.Println("Record data ======== ", recordData)
	err = ar.Conn.Create(&recordData).Error
	if err != nil {
		return detail_book.Core{}, err
	}
	return toCore1(recordData), nil
}

func (ar *mysqlDetailBookRepository) SelectAuthorbyBookID(id int) (authorID []int) {

	record := []DetailBook{}
	fmt.Println("id = ", id)
	if err := ar.Conn.Where("book_id = ?", id).Find(&record).Error; err != nil {
		return
	}
	fmt.Println(record)
	for _, value := range record {
		authorID = append(authorID, value.AuthorID)
	}
	return authorID
}
