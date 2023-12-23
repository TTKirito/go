postgres:
	docker compose up -d
create db: 
	docker exec -it postgres_container createdb --username=postgres --owner=postgres postgres
dropdb:
	docker exec -it postgres_container dropdb postgres
migrateup: 
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/postgres?sslmode=disable" --verbose up


migrateup1: 
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/postgres?sslmode=disable" --verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/postgres?sslmode=disable" --verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/postgres?sslmode=disable" --verbose down 1

sqlc: 
	sqlc generate

test: 
	go test -v -cover ./...

server:
	go run main.go

mock: 
	mockgen -package mockdb -destination db/mock/store.go github.com/TTKirito/go/db/sqlc Store
	
proto: 
	rm -f pb/*.go
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
    proto/*.proto

evans:
	evans --host localhost --port 9090 -r repl

.PHONY: postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 sqlc test server mock proto
