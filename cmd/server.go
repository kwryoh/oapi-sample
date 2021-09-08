package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"

	middleware "github.com/deepmap/oapi-codegen/pkg/chi-middleware"
	db "github.com/kwryoh/oapi-sample/db"
	api "github.com/kwryoh/oapi-sample/openapi"
)

var conn *sql.DB
var queries db.Queries
var ctx context.Context

const (
	dbdriver = "postgres"
	dbname = "example"
	dbpass = "pgpassword"
	dbuser = "postgres"
	dbhost = "db"
	dbport = "5432"
)

func main() {
	dbsource := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbhost, dbport, dbuser, dbpass, dbname, "disable"
	)
	conn, err = sql.Open(dbdriver, dbsource)
	if err != nil {
		fmt.Println()
	}
	defer conn.Close()

	ctx := context.Background()

	queries = db.New(conn)

	swagger, err := api.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading swagger spec¥n: %s", err)
		os.Exit(1)
	}

	swagger.Servers = nil

	itemStore := api.NewItemStore()

	r := chi.NewRouter()

	r.Use(middleware.OapiRequestValidator(swagger))

	api.HandlerFromMux(itemStore, r)

	addr := os.Getenv("Addr")
	if addr == "" {
		addr = ":9000"
	}

	log.Printf("Server listening on %s", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatalf("!! %+v", err)
	}
}
