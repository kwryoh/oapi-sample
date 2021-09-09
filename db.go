package main

import (
	"database/sql"

	_ "github.com/lib/pq"

	"github.com/kwryoh/oapi-sample/gen/db"
)

const (
	dbdriver = "postgres"
	dbsource = "host=db port=5432 user=postgres password=pgpostgres dbname=example sslmode=disale"
)

var queries *db.Queries

func ConnectDB() error {
	conn, err := sql.Open(dbdriver, dbsource)
	if err != nil {
		return err
	}
	defer conn.Close()

	queries = db.New(conn)

	return nil
}
