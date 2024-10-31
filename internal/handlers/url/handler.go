package url

import (
	"context"
)

type urlService interface {
	ShortenURL(ctx context.Context, originalURL, customAlias string) (string, error)
	GetOriginalURL(ctx context.Context, customAlias string) (string, error)
}

type URLHandler struct {
	urlService urlService
}

func NewURLHandler(urlService urlService) *URLHandler {
	return &URLHandler{
		urlService: urlService,
	}
}
