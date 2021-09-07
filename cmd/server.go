package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"

	middleware "github.com/deepmap/oapi-codegen/pkg/chi-middleware"
	api "github.com/kwryoh/oapi-sample/gen"
)

func main() {
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
