package db

import (
	"database/sql"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type Store interface {
	Querier
}

type SQLStore struct {
	db *sql.DB
	*Queries
}

func NewSQLStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

type DynamoDBStore struct {
	db        *dynamodb.Client
	TableName string
}

func NewDynamoDBStore(db *dynamodb.Client) Store {
	return &DynamoDBStore{
		db:        db,
		TableName: "Contracts",
	}
}
