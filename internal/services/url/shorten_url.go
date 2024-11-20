package url

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/nurmuh-alhakim18/url-shortener-api/internal/repositories"
)

func (s *URLService) ShortenURL(ctx context.Context, originalURL, customAlias string) (string, error) {
	exist, err := s.queries.CheckCustomAlias(ctx, customAlias)
	if err != nil {
		return "", err
	}

	if exist == 1 {
		return "", errors.New("custom alias is used")
	}

	validatedURL := validateAndPrependURL(originalURL)

	url, err := s.queries.CreateURL(ctx, repositories.CreateURLParams{
		CustomAlias:    customAlias,
		OriginalUrl:    validatedURL,
		ExpirationDate: sql.NullTime{Time: time.Now().UTC().Add(time.Hour * 24 * 7), Valid: true},
	})
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/%s", s.cfg.AppURL, url.CustomAlias), nil
}

func validateAndPrependURL(originalURL string) string {
	if !strings.HasPrefix(originalURL, "http://") && !strings.HasPrefix(originalURL, "https://") {
		return fmt.Sprintf("https://%s", originalURL)
	}

	return originalURL
}
