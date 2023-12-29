DB_URL=postgresql://root:secret@localhost:5432/postgres?sslmode=disable

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
	migrate -path db/migration -database "$(DB_URL)" --verbose down 1

new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)

sqlc: 
	sqlc generate

test: 
	go test -v -cover ./...

server:
	go run main.go

mock: 
	mockgen -package mockdb -destination db/mock/store.go github.com/TTKirito/go/db/sqlc Store
	mockgen -package mockwk -destination worker/mock/distributor.go github.com/TTKirito/go/worker TaskDistributor

proto: 
	rm -f pb/*.go
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
    --openapiv2_out=doc/swagger --openapiv2_opt=allow_merge=true,merge_file_name=simple_bank \
	proto/*.proto

evans:
	evans --host localhost --port 9090 -r repl

redis:
	docker run --name redis -p 6366:6379 -d redis:7-alpine

serverblog:
	go run streaming/blog_server/server.go


.PHONY: postgres createdb dropdb migrateup migratedown migrateup1 new_migration migratedown1 sqlc test server mock proto serverblog redis
