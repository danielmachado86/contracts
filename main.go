package main

import (
	"database/sql"

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

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		server.Logger.Fatalf("cannot connect to db:", err)
	}
	store := db.NewStore(conn)
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
