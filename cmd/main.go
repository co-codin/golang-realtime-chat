package main

import (
	"log"

	"github.com/co-codin/golang-realtime-chat/db"
)

func main() {
	_, err := db.NewDatabase()

	if err != nil {
		log.Fatalf("could not initialize database connection: %s", err)
	}

	
}