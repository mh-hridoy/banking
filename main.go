package main

import (
	"database/sql"
	"github/mh-hridoy/banking/api"
	db "github/mh-hridoy/banking/db/sqlc"
	"github/mh-hridoy/banking/utils"

	_ "github.com/lib/pq"

	"log"
)

func main() {
	config, errs := utils.LoadConfig(".")

	if errs != nil {
		log.Fatal("env cannot be loaded!")
	}
	// initialize server
	conn, err := sql.Open(config.DBDriver, config.DBurl)

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
