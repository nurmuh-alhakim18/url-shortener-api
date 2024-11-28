package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/nurmuh-alhakim18/gocache/cache"
	"github.com/nurmuh-alhakim18/url-shortener-api/config"
	handlerURL "github.com/nurmuh-alhakim18/url-shortener-api/internal/handlers/url"
	"github.com/nurmuh-alhakim18/url-shortener-api/internal/repositories"
	"github.com/nurmuh-alhakim18/url-shortener-api/internal/services/url"
	"github.com/nurmuh-alhakim18/url-shortener-api/router"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func main() {
	cfg := config.LoadConfig()

	cache := cache.NewCache(10)

	db, err := sql.Open("libsql", cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to make connection with db: %v", err)
	}

	queries := repositories.New(db)

	urlService := url.NewURLService(queries, cfg, cache)

	urlHandler := handlerURL.NewURLHandler(urlService)

	mux := router.NewRouter(urlHandler)

	server := http.Server{
		Addr:              ":" + cfg.Port,
		Handler:           mux,
		ReadHeaderTimeout: time.Second * 10,
	}

	log.Printf("Server running on port: %s\n", cfg.Port)
	log.Fatal(server.ListenAndServe())
}
