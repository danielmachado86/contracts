package main

import (
	"database/sql"
	"log"

	"github.com/danielmachado86/contracts/api"
	db "github.com/danielmachado86/contracts/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	serverAddress = "0.0.0.0:8080"
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:secret@localhost:5432/contracts?sslmode=disable"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
