package main

import (
	"log"

	"github.com/djangbahevans/go-template/api"
	"github.com/djangbahevans/go-template/config"
	"github.com/djangbahevans/go-template/databases/inmemory"
	"github.com/djangbahevans/go-template/services"
)

func main() {
	config.LoadConfig()

	db := inmemory.NewInMemoryDB()

	userService := services.NewUserService(db)
	server := api.NewApiServer(config.ServerAddr, userService)

	if err := server.Start(); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
