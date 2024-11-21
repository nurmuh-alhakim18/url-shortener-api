package router

import (
	"net/http"

	"github.com/nurmuh-alhakim18/url-shortener-api/internal/handlers/url"
	"github.com/nurmuh-alhakim18/url-shortener-api/pkg/utils"
)

func NewRouter(urlHandler *url.URLHandler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		utils.Response(w, http.StatusOK, "OK")
	})

	mux.HandleFunc("POST /api/shorten", urlHandler.HandlerShortenURL)
	mux.HandleFunc("GET /{alias}", urlHandler.HandlerRedirectURL)

	return mux
}
