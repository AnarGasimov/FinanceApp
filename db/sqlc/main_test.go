package db

import (
	"FinanceApp/config"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
	"testing"
)

const (
	dbDriver = "postgres"
)

var testQueries *Queries
var testStore Store

func TestMain(m *testing.M) {

	cfg, err := config.MustLoad()

	if err != nil {
		log.Fatal("Failed to load Configuration {}", err)
	}

	conn, err := sql.Open(dbDriver, cfg.Database.URL)

	if err != nil {
		log.Fatal("cannot connect to DB {}", err)
	}

	testQueries = New(conn)
	testStore = *NewStore(conn)
	os.Exit(m.Run())
}
