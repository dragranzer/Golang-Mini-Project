package main

import (
	"fmt"
	"os"

	"github.com/dragranzer/Golang-Mini-Project/config"
	"github.com/dragranzer/Golang-Mini-Project/migrate"
	"github.com/dragranzer/Golang-Mini-Project/routes"
)

func main() {

	config.InitDB()
	migrate.AutoMigrate()

	e := routes.New()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	e.Start(":" + port)
	fmt.Println("Server started in port :1234")
}
