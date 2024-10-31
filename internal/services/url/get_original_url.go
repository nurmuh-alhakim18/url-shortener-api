package url

import (
	"context"
	"errors"
)

func (s *URLService) GetOriginalURL(ctx context.Context, customAlias string) (string, error) {
	originalUrl, err := s.queries.GetOriginalURL(ctx, customAlias)
	if err != nil {
		return "", errors.New("failed to get orinal url")
	}

	return originalUrl, nil
}
