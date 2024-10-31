package url

import (
	"encoding/json"
	"net/http"

	"github.com/nurmuh-alhakim18/url-shortener-api/internal/models/url"
	"github.com/nurmuh-alhakim18/url-shortener-api/pkg/utils"
)

func (h *URLHandler) HandlerShortenURL(w http.ResponseWriter, r *http.Request) {
	var params url.ShortenURLReq
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		utils.ResponseError(w, http.StatusBadRequest, "invalid input", err)
		return
	}

	if params.CustomAlias == "" || params.OriginalUrl == "" {
		utils.ResponseError(w, http.StatusBadRequest, "custom alias or original url have to be filled", err)
		return
	}

	generatedURL, err := h.urlService.ShortenURL(r.Context(), params.OriginalUrl, params.CustomAlias)
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err.Error(), err)
		return
	}

	utils.Response(w, http.StatusCreated, url.ShortenURLResp{
		GeneratedLink: generatedURL,
	})
}
