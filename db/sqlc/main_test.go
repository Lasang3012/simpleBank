package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://postgres:123456@localhost:5432/postgres?sslmode=disable"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error

	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}

// var err error
// 	config, err := util.LoadConfig("../..")
// 	if err != nil {
// 		log.Fatal("cannot load config:", err)
// 	}

// 	testDB, err = sql.Open(config.DBDriver, config.DBSource)
// 	if err != nil {
// 		log.Fatal("cannot connect to db:", err)
// 	}

// 	testQueries = New(testDB)

// 	os.Exit(m.Run())
