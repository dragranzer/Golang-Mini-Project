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
	return toCore1(recordData), nil
}

func (ar *mysqlBookRepository) SelectAllData() (resp []books.Core) {
	record := []Book{}

	if err := ar.Conn.Preload("Kategoris").Find(&record).Error; err != nil {
		return []books.Core{}
	}

	// jangan lupa ditranlasiin ke core
	return toCoreList(record)
}

func (br *mysqlBookRepository) SelectData(judul string) (resp []books.Core, err error) {

	record := []Book{}
	// fmt.Println("judul = ", judul)
	if err = br.Conn.Preload("Kategoris").Where("judul = ?", judul).Find(&record).Error; err != nil {
		return []books.Core{}, err
	}

	return toCoreList(record), nil
}

func (br *mysqlBookRepository) UpdateTersediabyName(name string) (resp books.Core, err error) {
	record := Book{}
	tersedia := Book{}
	if err = br.Conn.Preload("Kategoris").Where("judul = ?", name).Find(&tersedia).Error; err != nil {
		return books.Core{}, err
	}
	tersedia.Tersedia--
	// fmt.Println("tersedia ", tersedia)
	if err = br.Conn.Model(&record).Where("id = ?", tersedia.ID).Update("tersedia", tersedia.Tersedia).Error; err != nil {
		return books.Core{}, err
	}
	// fmt.Println("update ", record)
	return
}
