package url

import (
	"context"
	"time"

	"github.com/nurmuh-alhakim18/url-shortener-api/config"
	"github.com/nurmuh-alhakim18/url-shortener-api/internal/repositories"
)

type urlRepository interface {
	CheckCustomAlias(ctx context.Context, customAlias string) (int64, error)
	CreateURL(ctx context.Context, arg repositories.CreateURLParams) (repositories.Url, error)
	GetOriginalURL(ctx context.Context, customAlias string) (string, error)
}

type cacheInterface interface {
	Get(key string) (interface{}, bool)
	Set(key string, value interface{}, ttl time.Duration)
}

type URLService struct {
	queries urlRepository
	cfg     config.Config
	cache   cacheInterface
}

func NewURLService(urlRepo urlRepository, cfg config.Config, cache cacheInterface) *URLService {
	return &URLService{
		queries: urlRepo,
		cfg:     cfg,
		cache:   cache,
	}
}
