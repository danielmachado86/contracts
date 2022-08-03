package main

import (
	"github.com/danielmachado86/contracts/api"
	db "github.com/danielmachado86/contracts/db"
	"github.com/danielmachado86/contracts/utils"
	_ "github.com/lib/pq"
)

func main() {

	server, err := api.NewServer()
	server.Logger.Infof("creating server...")
	if err != nil {
		server.Logger.Fatalf("cannot create server:", err)
	}

	config, err := utils.LoadConfig("./.env")
	if err != nil {
		server.Logger.Fatalf("cannot load config:", err)
	}
	dynamoClient, err := db.CreateLocalClient()
	store := db.NewDynamoDBStore(dynamoClient)

	server.Logger.Infof("configuring server...")
	err = server.ConfigServer(config, store)
	if err != nil {
		server.Logger.Fatalf("cannot configure server", err)
	}
	server.Logger.Infof("server listening for requests...")
	err = server.Start(config.ServerAddress)
	if err != nil {
		server.Logger.Fatalf("cannot start server:", err)
	}

}
