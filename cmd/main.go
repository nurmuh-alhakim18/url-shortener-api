package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"
	"github.com/nurmuh-alhakim18/url-shortener-api/config"
	handlerURL "github.com/nurmuh-alhakim18/url-shortener-api/internal/handlers/url"
	"github.com/nurmuh-alhakim18/url-shortener-api/internal/repositories"
	"github.com/nurmuh-alhakim18/url-shortener-api/internal/services/url"
	"github.com/nurmuh-alhakim18/url-shortener-api/router"
)

func main() {
	cfg := config.LoadConfig()

	db, err := sql.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to make connection with db: %v", err)
	}

	queries := repositories.New(db)

	urlService := url.NewURLService(queries, cfg)

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
