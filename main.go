package main

import (
	"database/sql"
	"github/mh-hridoy/banking/api"
	db "github/mh-hridoy/banking/db/sqlc"

	_ "github.com/lib/pq"

	"log"
)

// connect db

var (
	sqlDriver = "postgres"
	sqlUrl    = "postgresql://hridoy:2543@localhost:5432/go1?sslmode=disable"
)

func main() {
	// initialize server
	conn, err := sql.Open(sqlDriver, sqlUrl)

	if err != nil {
		log.Fatal("cannot connect to db", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.StartServer()

	if err != nil {
		log.Fatal(err)
	}

}
