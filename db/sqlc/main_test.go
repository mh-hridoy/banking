package db

import (
	"database/sql"
	"github/mh-hridoy/banking/utils"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDb *sql.DB

func TestMain(m *testing.M) {
	config, errs := utils.LoadConfig("../..")

	if errs != nil {
		log.Fatal("env cannot be loaded!")
	}
	var err error
	testDb, err = sql.Open(config.DBDriver, config.DBurl)

	if err != nil {
		log.Fatal("cannot connect to db", err)
	}

	testQueries = New(testDb)

	os.Exit(m.Run())

}
