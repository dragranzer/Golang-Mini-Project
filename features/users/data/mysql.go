package data

import (
	"fmt"

	"github.com/dragranzer/Golang-Mini-Project/features/users"
	"gorm.io/gorm"
)

type mysqlUserRepository struct {
	Conn *gorm.DB
}

func NewUserRepository(conn *gorm.DB) users.Data {
	return &mysqlUserRepository{
		Conn: conn,
	}
}

func (ar *mysqlUserRepository) InsertData(data users.Core) (resp users.Core, err error) {
	// fmt.Println("Data ========", data)
	recordData := fromCore(data)
	fmt.Println("Record data ======== ", recordData)
	err = ar.Conn.Create(&recordData).Error
	if err != nil {
		return users.Core{}, err
	}
	return data, nil
}

func (ar *mysqlUserRepository) SelectAllData() (resp []users.Core) {
	record := []User{}

	if err := ar.Conn.Find(&record).Error; err != nil {
		return []users.Core{}
	}

	// jangan lupa ditranlasiin ke core
	return toCoreList(record)
}
