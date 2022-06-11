package main

import (
	"database/sql"
	"log"

	"github.com/danielmachado86/contracts/api"
	db "github.com/danielmachado86/contracts/db/sqlc"
	"github.com/danielmachado86/contracts/utils"
	_ "github.com/lib/pq"
)

func main() {
	config, err := utils.LoadConfig("./.env")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	store := db.NewStore(conn)
	server, err := api.NewServer(store)
	if err != nil {
		log.Fatal("cannot create server", err)
	}
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
