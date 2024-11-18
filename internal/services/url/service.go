package url

import (
	"context"
	"time"

	"github.com/nurmuh-alhakim18/url-shortener-api/config"
	"github.com/nurmuh-alhakim18/url-shortener-api/internal/repositories"
)

type urlRepository interface {
	CheckCustomAlias(ctx context.Context, customAlias string) (bool, error)
	CreateURL(ctx context.Context, arg repositories.CreateURLParams) (repositories.Url, error)
	GetOriginalURL(ctx context.Context, customAlias string) (string, error)
}

type redisInterface interface {
	Set(ctx context.Context, key, value string, ttl time.Duration) error
	Get(ctx context.Context, key string) (string, error)
}

type URLService struct {
	queries urlRepository
	cfg     config.Config
	redis   redisInterface
}

func NewURLService(urlRepo urlRepository, cfg config.Config, redis redisInterface) *URLService {
	return &URLService{
		queries: urlRepo,
		cfg:     cfg,
		redis:   redis,
	}
}
