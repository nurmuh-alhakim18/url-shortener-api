package router

import (
	"net/http"

	"github.com/nurmuh-alhakim18/url-shortener-api/internal/handlers/url"
)

func NewRouter(urlHandler *url.URLHandler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /api/shorten", urlHandler.HandlerShortenURL)
	mux.HandleFunc("GET /{alias}", urlHandler.HandlerRedirectURL)

	return mux
}
