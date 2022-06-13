package main

import (
	"log"
	"read/app"
	"read/config"
	"read/db"
)

func main() {
	config.LoadConfig()

	db, err := db.NewDatabase()
	if err != nil {
		log.Fatal(err.Error())
	}

	server, err := app.NewServer(db)
	if err != nil {
		log.Fatal(err.Error())
	}

	server.StartServer("5002")
}