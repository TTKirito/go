package main

import (
	"database/sql"
	"log"

	"github.com/TTKirito/go/api"
	db "github.com/TTKirito/go/db/sqlc"
	"github.com/TTKirito/go/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")

	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSOURCE)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)

	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start((config.ServerAddress))

	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
