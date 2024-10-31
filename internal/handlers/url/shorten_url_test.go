package url

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nurmuh-alhakim18/url-shortener-api/internal/models/url"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func (m *MockURLService) ShortenURL(ctx context.Context, originalURL, customAlias string) (string, error) {
	args := m.Called(ctx, originalURL, customAlias)
	return args.String(0), args.Error(1)
}

func TestURLHandler_HandlerShortenURL(t *testing.T) {
	tests := []struct {
		name       string
		reqBody    *url.ShortenURLReq
		mockFn     func(m *MockURLService)
		wantStatus int
		wantResp   *url.ShortenURLResp
	}{
		{
			name: "Successful URL shortening",
			reqBody: &url.ShortenURLReq{
				OriginalUrl: "https://original-url.com",
				CustomAlias: "alias",
			},
			mockFn: func(m *MockURLService) {
				m.On("ShortenURL", mock.Anything, "https://original-url.com", "alias").Return("http://localhost:8080/alias", nil)
			},
			wantStatus: http.StatusCreated,
			wantResp: &url.ShortenURLResp{
				GeneratedLink: "http://localhost:8080/alias",
			},
		},
		{
			name: "Empty custom alias",
			reqBody: &url.ShortenURLReq{
				OriginalUrl: "https://original-url.com",
				CustomAlias: "",
			},
			mockFn: func(m *MockURLService) {
			},
			wantStatus: http.StatusBadRequest,
			wantResp:   nil,
		},
		{
			name: "Empty original url",
			reqBody: &url.ShortenURLReq{
				OriginalUrl: "",
				CustomAlias: "alias",
			},
			mockFn:     func(m *MockURLService) {},
			wantStatus: http.StatusBadRequest,
			wantResp:   nil,
		},
		{
			name: "Shorten URL service error",
			reqBody: &url.ShortenURLReq{
				OriginalUrl: "https://original-url.com",
				CustomAlias: "alias",
			},
			mockFn: func(m *MockURLService) {
				m.On("ShortenURL", mock.Anything, "https://original-url.com", "alias").Return("", assert.AnError)
			},
			wantStatus: http.StatusInternalServerError,
			wantResp:   nil,
		},
		{
			name:       "Invalid JSON request body",
			reqBody:    nil,
			mockFn:     func(m *MockURLService) {},
			wantStatus: http.StatusBadRequest,
			wantResp:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockURLService := &MockURLService{}
			tt.mockFn(mockURLService)

			h := &URLHandler{
				urlService: mockURLService,
			}

			data, err := json.Marshal(tt.reqBody)
			if err != nil {
				assert.Error(t, err)
			}

			assert.NoError(t, err)

			req := httptest.NewRequest("GET", "/api/shorten", bytes.NewBuffer(data))
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()

			mux := http.NewServeMux()
			mux.HandleFunc("/api/shorten", h.HandlerShortenURL)

			h.HandlerShortenURL(rec, req)

			assert.Equal(t, tt.wantStatus, rec.Code)

			if tt.wantResp != nil {
				var gotResp url.ShortenURLResp
				err = json.NewDecoder(rec.Body).Decode(&gotResp)
				assert.NoError(t, err)
				assert.Equal(t, tt.wantResp.GeneratedLink, gotResp.GeneratedLink)
			}

			mockURLService.AssertExpectations(t)
		})
	}
}
