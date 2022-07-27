package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver        = "postgres"
	dbSource        = "postgresql://root:secret@localhost:5432/contracts?sslmode=disable"
	dbFailingSource = "failing Source"
)

var testQueries *Queries
var testFailingQueries *Queries
var testDB *sql.DB
var testFailingDB *sql.DB

func TestMain(m *testing.M) {

	testFailingDB, _ = sql.Open(dbDriver, dbFailingSource)
	testFailingQueries = New(testFailingDB)

	var err error

	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
