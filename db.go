package main

import (
	"context"
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

var (
	conn    *sql.DB
	queries *db.Queries
	ctx     context.Context
)

func ConnectDB() error {
	source := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbname)
	log.Printf("source: %s", source)

	var err error
	conn, err = sql.Open(driver, source)
	if err != nil {
		return err
	}

	log.Printf("Successfully connected.")
	ctx = context.Background()
	queries = db.New(conn)

	return nil
}
