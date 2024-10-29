package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/nurmuh-alhakim18/url-shortener-api/config"
	"github.com/nurmuh-alhakim18/url-shortener-api/router"
)

func main() {
	cfg := config.LoadConfig()

	db, err := sql.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to make connection with db: %v", err)
	}

	mux := router.NewRouter()

	server := http.Server{
		Addr:    ":" + cfg.Port,
		Handler: mux,
	}

	log.Printf("Server running on port: %s\n", cfg.Port)
	log.Fatal(server.ListenAndServe())
}
