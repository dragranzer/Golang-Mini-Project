package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {

	var err error
	db, err := gorm.Open(mysql.Open("sql6456325:alIJAGMBRw@tcp(sql6.freemysqlhosting.net:3306)/sql6456325?parseTime=true"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB = db
}
