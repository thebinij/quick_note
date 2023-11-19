swaggerGen:
	swagger generate spec -o ./swaggerui/swagger.json

postgresinit:
	docker run --name postgres15 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:15-alpine

postgres:
	docker exec -it postgres15 psql

createdb:
	docker exec -it postgres15 createdb --username=root --owner=root quick-note

dropdb:
	docker exec -it postgres15 dropdb quick-note

migrateup:
	migrate -path db/migrations -database "postgresql://root:password@localhost:5432/quick-note?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://root:password@localhost:5432/quick-note?sslmode=disable" -verbose down

.PHONY: postgresinit postgres createdb dropdb migrateup migratedown swagger