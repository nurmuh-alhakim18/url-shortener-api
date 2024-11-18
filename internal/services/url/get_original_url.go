package url

import (
	"context"
	"errors"
	"time"
)

func (s *URLService) GetOriginalURL(ctx context.Context, customAlias string) (string, error) {
	originalURL, err := s.redis.Get(ctx, customAlias)
	if err == nil {
		return originalURL, nil
	}

	originalURL, err = s.queries.GetOriginalURL(ctx, customAlias)
	if err != nil {
		return "", errors.New("failed to get orinal url")
	}

	_ = s.redis.Set(ctx, customAlias, originalURL, time.Hour)

	return originalURL, nil
}
