package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	chi_middleware "github.com/go-chi/chi/v5/middleware"

	middleware "github.com/deepmap/oapi-codegen/pkg/chi-middleware"
	app "github.com/kwryoh/oapi-sample"
	api "github.com/kwryoh/oapi-sample/gen/openapi"
)

func main() {
	var conn sql.DB
	queries, err := app.ConnectDB(&conn)
	if err != nil {
		log.Fatal("Cannot connect database: ", err)
		os.Exit(-1)
	}
	defer conn.Close()

	swagger, err := api.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading swagger spec¥n: %s", err)
		os.Exit(1)
	}
	swagger.Servers = nil

	ctx := context.Background()
	itemStore := app.NewItemStore(queries, ctx)

	r := chi.NewRouter()
	r.Use(middleware.OapiRequestValidator(swagger))
	r.Use(chi_middleware.RealIP)
	r.Use(chi_middleware.Logger)
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
