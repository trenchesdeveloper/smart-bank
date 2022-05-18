postgres:
	docker run -d -p 5432:5432 -e POSTGRES_PASSWORD=secret --name postgres12 postgres:12-alpine

createdb:
	docker exec -it postgres12  createdb -U postgres smartBank
	
dropdb:
	docker exec -it postgres12 dropdb smartBank
migrateup:
	migrate -path db/migration -database "postgresql://postgres:secret@localhost:5432/smartBank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:secret@localhost:5432/smartBank?sslmode=disable" -verbose down



sqlc:
	sqlc generate

.PHONY: postgres createdb dropdb migrateup migratedown sqlc