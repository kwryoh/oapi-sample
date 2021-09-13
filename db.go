package app

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/kwryoh/oapi-sample/gen/db"
)

const (
	driver   = "postgres"
	host     = "db"
	port     = "5432"
	user     = "postgres"
	password = "pgpassword"
	dbname   = "example"
)

func ConnectDB(conn *sql.DB) (*db.Queries, error) {
	var queries *db.Queries
	source := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbname)
	log.Printf("source: %s", source)

	var err error
	conn, err = sql.Open(driver, source)
	if err != nil {
		return queries, err
	}

	log.Printf("Successfully connected.")
	queries = db.New(conn)

	return queries, nil
}
