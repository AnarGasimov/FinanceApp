# Makefile
postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=postgres -d postgres

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root finance

dropdb:
	docker exec -it postgres12 psql -U root -c "SELECT pg_terminate_backend(pid) FROM pg_stat_activity WHERE datname = 'finance' AND pid <> pg_backend_pid();"
	docker exec -it postgres12 dropdb finance

migrateup:
	migrate -path db/migration -database "postgres://root:postgres@localhost:5432/finance?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgres://root:postgres@localhost:5432/finance?sslmode=disable" -verbose down

sqlc:
	sqlc generate
test:
	go test -v -cover ./...
.PHONY: postgres createdb dropdb migrateup migratedown sqlc test
