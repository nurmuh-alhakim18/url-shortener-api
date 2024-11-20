package url

import (
	"context"

	"github.com/nurmuh-alhakim18/url-shortener-api/config"
	"github.com/nurmuh-alhakim18/url-shortener-api/internal/repositories"
)

type urlRepository interface {
	CheckCustomAlias(ctx context.Context, customAlias string) (int64, error)
	CreateURL(ctx context.Context, arg repositories.CreateURLParams) (repositories.Url, error)
	GetOriginalURL(ctx context.Context, customAlias string) (string, error)
}

type URLService struct {
	queries urlRepository
	cfg     config.Config
}

func NewURLService(urlRepo urlRepository, cfg config.Config) *URLService {
	return &URLService{
		queries: urlRepo,
		cfg:     cfg,
	}
}
