package url

import (
	"context"
	"time"
)

func (s *URLService) GetOriginalURL(ctx context.Context, customAlias string) (string, error) {
	if originalUrl, ok := s.cache.Get(customAlias); ok {
		return originalUrl.(string), nil
	}

	originalUrl, err := s.queries.GetOriginalURL(ctx, customAlias)
	if err != nil {
		return "", err
	}

	s.cache.Set(customAlias, originalUrl, time.Hour)

	return originalUrl, nil
}
