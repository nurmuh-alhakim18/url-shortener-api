package router

import (
	"net/http"

	"github.com/nurmuh-alhakim18/url-shortener-api/internal/handlers/url"
)

func NewRouter(urlHandler *url.URLHandler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/shorten", urlHandler.HandlerShortenURL)
	mux.HandleFunc("/{alias}", urlHandler.HandlerRedirectURL)

	return mux
}
