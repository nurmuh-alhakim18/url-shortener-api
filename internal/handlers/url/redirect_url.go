package url

import (
	"net/http"

	"github.com/nurmuh-alhakim18/url-shortener-api/pkg/utils"
)

func (h *URLHandler) HandlerRedirectURL(w http.ResponseWriter, r *http.Request) {
	customAlias := r.PathValue("alias")
	originalURL, err := h.urlService.GetOriginalURL(r.Context(), customAlias)
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err.Error(), err)
		return
	}

	http.Redirect(w, r, originalURL, http.StatusFound)
}
