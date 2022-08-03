DB_URL=postgresql://root:secret@localhost:5432/contracts?sslmode=disable

postgres:
	docker run --name postgres --network contracts-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14.2-alpine

createdb:
	docker exec -it postgres createdb --username=root --owner=root contracts

dropdb:
	docker exec -it postgres dropdb contracts

migrateup:
	migrate -path db/migrations -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migrations -database "$(DB_URL)" -verbose down

migrateup1:
	migrate -path db/migrations -database "$(DB_URL)" -verbose up 1

migratedown1:
	migrate -path db/migrations -database "$(DB_URL)" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/danielmachado86/contracts/db Store

.PHONY: postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 sqlc test server mock
