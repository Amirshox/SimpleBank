postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=amirshokh -e POSTGRES_PASSWORD=secret -d postgres:14-alpine

createdb:
	docker exec -it postgres createdb --username=amirshokh --owner=amirshokh simple_bank

dropdb:
	docker exec -it postgres dropdb simple_bank -U amirshokh

migrateup:
	migrate -path db/migration -database "postgresql://amirshokh:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://amirshokh:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY: postgres createdb dropdb migrateup migratedown sqlc