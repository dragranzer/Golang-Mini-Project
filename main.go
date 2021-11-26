package main

import (
	"fmt"

	"github.com/dragranzer/Golang-Mini-Project/config"
	"github.com/dragranzer/Golang-Mini-Project/migrate"
	"github.com/dragranzer/Golang-Mini-Project/routes"
)

func main() {

	config.InitDB()
	migrate.AutoMigrate()

	e := routes.New()

	e.Start(":1234")
	fmt.Println("Server started in port :1234")
}
