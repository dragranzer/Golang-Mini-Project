package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {

	var err error
	db, err := gorm.Open(mysql.Open("root:@tcp(host.docker.internal:3306)/minpro_be?parseTime=true"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB = db
}
