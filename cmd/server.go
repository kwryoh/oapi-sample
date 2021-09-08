package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"

	middleware "github.com/deepmap/oapi-codegen/pkg/chi-middleware"
	api "github.com/kwryoh/oapi-sample/openapi"
)

var db *sql.DB

const (
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
	db, err := sql.Open("postgres", dbsource)
	defer db.Close()
	if err != nil {
		fmt.Println()
	}


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
