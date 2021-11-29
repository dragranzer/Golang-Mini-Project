package data

import (
	"github.com/dragranzer/Golang-Mini-Project/features/peminjamans"
	"gorm.io/gorm"
)

type mysqlPeminjamanRepository struct {
	Conn *gorm.DB
}

func NewPeminjamanRepository(conn *gorm.DB) peminjamans.Data {
	return &mysqlPeminjamanRepository{
		Conn: conn,
	}
}

func (ar *mysqlPeminjamanRepository) InsertData(data peminjamans.Core) (resp peminjamans.Core, err error) {
	// fmt.Println("Data ========", data)
	recordData := fromCore(data)
	// fmt.Println("Record data ======== ", recordData)
	err = ar.Conn.Create(&recordData).Error
	if err != nil {
		return peminjamans.Core{}, err
	}
	return data, nil
}

func (ar *mysqlPeminjamanRepository) SelectAllData() (resp []peminjamans.Core) {
	record := []Peminjaman{}

	if err := ar.Conn.Find(&record).Error; err != nil {
		return []peminjamans.Core{}
	}

	// jangan lupa ditranlasiin ke core
	return toCoreList(record)
}
