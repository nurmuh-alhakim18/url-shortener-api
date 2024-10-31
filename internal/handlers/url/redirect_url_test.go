package url

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockURLService struct {
	mock.Mock
}

func (m *MockURLService) GetOriginalURL(ctx context.Context, customAlias string) (string, error) {
	args := m.Called(ctx, customAlias)
	return args.String(0), args.Error(1)
}

func TestURLHandler_HandlerRedirectURL(t *testing.T) {
	tests := []struct {
		name       string
		alias      string
		mockFn     func(m *MockURLService)
		wantStatus int
		wantURL    string
	}{
		{
			name:  "Successful redirect URL",
			alias: "alias",
			mockFn: func(m *MockURLService) {
				m.On("GetOriginalURL", mock.Anything, "alias").Return("https://original-url.com", nil)
			},
			wantStatus: http.StatusFound,
			wantURL:    "https://original-url.com",
		},
		{
			name:  "Get original URL service error",
			alias: "invalidAlias",
			mockFn: func(m *MockURLService) {
				m.On("GetOriginalURL", mock.Anything, "invalidAlias").Return("", assert.AnError)
			},
			wantStatus: http.StatusInternalServerError,
			wantURL:    "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockURLService := &MockURLService{}
			tt.mockFn(mockURLService)

			h := &URLHandler{
				urlService: mockURLService,
			}

			req := httptest.NewRequest(http.MethodGet, "/"+tt.alias, nil)
			rec := httptest.NewRecorder()

			mux := http.NewServeMux()
			mux.HandleFunc("/{alias}", h.HandlerRedirectURL)
			mux.ServeHTTP(rec, req)

			assert.Equal(t, tt.wantStatus, rec.Code)

			if tt.wantURL != "" {
				assert.Equal(t, tt.wantURL, rec.Header().Get("Location"))
			}

			mockURLService.AssertExpectations(t)
		})
	}
}
