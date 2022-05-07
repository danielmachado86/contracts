DB_URL=postgresql://root:secret@host.docker.internal:5432/contracts?sslmode=disable

postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14.2-alpine

createdb:
	docker exec -it postgres createdb --username=root --owner=root contracts

dropdb:
	docker exec -it postgres dropdb contracts

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down
