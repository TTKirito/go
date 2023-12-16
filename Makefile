postgres:
	docker compose up -d
create db: 
	docker exec -it postgres_container createdb --username=postgres --owner=postgres postgres
dropdb:
	docker exec -it postgres_container dropdb postgres
migrateup: 
	migrate -path db/migration -database "postgres://postgres:changeme@localhost:5433/postgres?sslmode=disable" --verbose up
migratedown:
	migrate -path db/migration -database "postgres://postgres:changeme@localhost:5433/postgres?sslmode=disable" --verbose down
sqlc: 
	sqlc generate

test: 
	go test -v -cover ./...

server:
	go run main.go

mock: 
	mockgen -package mockdb -destination db/mock/store.go github.com/TTKirito/go/db/sqlc Store
	
.PHONY: postgres createdb dropdb migrateup migratedown sqlc
