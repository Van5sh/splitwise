package main

import (
	"log"

	database "github.com/Van5sh/new-splitwise/internal/db"
)

func main() {
	db, err := database.ConnectToDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	startServer(db)
}
