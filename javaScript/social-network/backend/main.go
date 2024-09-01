package main

import (
	"log"
	"social-network/data/database"
	"social-network/pkg/server"
)

func main() {
	db, err := database.OpenDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	server.Start(db)
}
