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

func (ar *mysqlUserRepository) InsertData(data users.Core) (err error) {
	// fmt.Println("Data ========", data)
	recordData := fromCore(data)
	fmt.Println("Record data ======== ", recordData)
	err = ar.Conn.Create(&recordData).Error
	if err != nil {
		return err
	}
	return nil
}

func (ar *mysqlUserRepository) SelectAllData() (resp []users.Core) {
	record := []User{}

	if err := ar.Conn.Find(&record).Error; err != nil {
		return []users.Core{}
	}

	// jangan lupa ditranlasiin ke core
	return toCoreList(record)
}

func (ar *mysqlUserRepository) SelectDatabyName(name string) (resp users.Core, err error) {

	record := User{}
	if err = ar.Conn.Where("nama = ?", name).Find(&record).Error; err != nil {
		return users.Core{}, err
	}

	return record.toCore(), nil
}

func (ar *mysqlUserRepository) SelectDatabyID(id int) (resp users.Core, err error) {
	record := User{}
	if err = ar.Conn.Where("id = ?", id).Find(&record).Error; err != nil {
		return users.Core{}, err
	}

	return record.toCore(), nil
}

func (ar *mysqlUserRepository) UpdateDatabyID(id int, newData users.Core) (err error) {
	// db.Model(&User{}).Where("active = ?", true).Update("name", "hello")
	record := User{
		Nama:  newData.Nama,
		Email: newData.Email,
	}
	if err = ar.Conn.Model(&record).Where("id = ?", id).Updates(User{Nama: record.Nama, Email: newData.Email}).Error; err != nil {
		return err
	}
	return
}

func (ar *mysqlUserRepository) CheckEmailPass(email string, pass string) (isAuth bool, user users.Core, err error) {
	record := User{}
	err = ar.Conn.Where("email = ? AND password = ?", email, pass).First(&record).Error
	if err != nil {
		return false, user, err
	}
	if record.ID == 0 {
		return false, user, nil
	}
	return true, record.toCore(), nil
}

func (ar *mysqlUserRepository) SelectDatabyEmail(email string) (resp users.Core, err error) {
	record := User{}
	if err = ar.Conn.Where("email = ?", email).Find(&record).Error; err != nil {
		return users.Core{}, err
	}

	return record.toCore(), nil
}
