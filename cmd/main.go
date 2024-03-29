package main

import (
	"log"

	"github.com/co-codin/golang-realtime-chat/db"
	"github.com/co-codin/golang-realtime-chat/internal/user"
	"github.com/co-codin/golang-realtime-chat/internal/ws"
	"github.com/co-codin/golang-realtime-chat/router"
)

func main() {
	dbConn, err := db.NewDatabase()

	if err != nil {
		log.Fatalf("could not initialize database connection: %s", err)
	}

	userRep := user.NewRepository(dbConn.GetDB())
	userSvc := user.NewService(userRep)
	userHandler := user.NewHandler(userSvc)

	hub := ws.NewHub()
	wsHandler := ws.NewHandler(hub)

	go hub.Run()


	router.InitRouter(userHandler, wsHandler)
	router.Start("0.0.0.0:8080")
}