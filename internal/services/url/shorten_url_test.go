package url

import (
	"context"
	"testing"

	"github.com/nurmuh-alhakim18/url-shortener-api/config"
	"github.com/nurmuh-alhakim18/url-shortener-api/internal/repositories"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockQueries struct {
	mock.Mock
}

func (m *MockQueries) CheckCustomAlias(ctx context.Context, customAlias string) (bool, error) {
	args := m.Called(ctx, customAlias)
	return args.Bool(0), args.Error(1)
}

func (m *MockQueries) CreateURL(ctx context.Context, arg repositories.CreateURLParams) (repositories.Url, error) {
	args := m.Called(ctx, arg)
	return args.Get(0).(repositories.Url), args.Error(1)
}

func TestURLService_ShortenURL(t *testing.T) {
	cfg := config.Config{AppURL: "http://localhost:8080"}

	type args struct {
		ctx         context.Context
		originalURL string
		customAlias string
	}

	tests := []struct {
		name    string
		args    args
		mockFn  func(m *MockQueries)
		want    string
		wantErr bool
	}{
		{
			name: "Custom alias exists",
			args: args{
				ctx:         context.Background(),
				originalURL: "https://original-url.com",
				customAlias: "alias",
			},
			mockFn: func(m *MockQueries) {
				m.On("CheckCustomAlias", mock.Anything, "alias").Return(true, nil)
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Successful URL shortening",
			args: args{
				ctx:         context.Background(),
				originalURL: "https://original-url.com",
				customAlias: "alias",
			},
			mockFn: func(m *MockQueries) {
				m.On("CheckCustomAlias", mock.Anything, "alias").Return(false, nil)
				m.On("CreateURL", mock.Anything, mock.AnythingOfType("repositories.CreateURLParams")).Return(repositories.Url{CustomAlias: "alias"}, nil)
			},
			want:    "http://localhost:8080/alias",
			wantErr: false,
		},
		{
			name: "URL without protocol - should prepend https",
			args: args{
				ctx:         context.Background(),
				originalURL: "original-url.com",
				customAlias: "alias",
			},
			mockFn: func(m *MockQueries) {
				m.On("CheckCustomAlias", mock.Anything, "alias").Return(false, nil)
				m.On("CreateURL", mock.Anything, mock.AnythingOfType("repositories.CreateURLParams")).Return(repositories.Url{CustomAlias: "alias"}, nil)
			},
			want:    "http://localhost:8080/alias",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockQueries := &MockQueries{}
			tt.mockFn(mockQueries)

			s := &URLService{
				queries: mockQueries,
				cfg:     cfg,
			}

			got, err := s.ShortenURL(tt.args.ctx, tt.args.originalURL, tt.args.customAlias)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}

			mockQueries.AssertExpectations(t)
		})
	}
}
