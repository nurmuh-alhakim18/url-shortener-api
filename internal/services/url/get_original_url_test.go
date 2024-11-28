package url

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/nurmuh-alhakim18/url-shortener-api/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func (m *MockQueries) GetOriginalURL(ctx context.Context, customAlias string) (string, error) {
	args := m.Called(ctx, customAlias)
	return args.String(0), args.Error(1)
}

type MockCache struct {
	mock.Mock
}

func (m *MockCache) Get(key string) (interface{}, bool) {
	args := m.Called(key)
	return args.Get(0), args.Bool(1)
}

func (m *MockCache) Set(key string, value interface{}, ttl time.Duration) {
	m.Called(key, value, ttl)
}

func TestURLService_GetOriginalURL(t *testing.T) {
	cfg := config.Config{}
	type args struct {
		ctx         context.Context
		customAlias string
	}
	tests := []struct {
		name        string
		args        args
		mockFn      func(m *MockQueries)
		mockFnCache func(m *MockCache)
		want        string
		wantErr     bool
	}{
		{
			name: "Cache hit for original URL",
			args: args{
				ctx:         context.Background(),
				customAlias: "alias",
			},
			mockFn: func(m *MockQueries) {
			},
			mockFnCache: func(m *MockCache) {
				m.On("Get", "alias").Return("http://kimoutofkim.blog", true)
			},
			want:    "http://kimoutofkim.blog",
			wantErr: false,
		},
		{
			name: "Cache miss, database hit for original URL",
			args: args{
				ctx:         context.Background(),
				customAlias: "alias",
			},
			mockFn: func(m *MockQueries) {
				m.On("GetOriginalURL", mock.Anything, "alias").Return("http://kimoutofkim.blog", nil)
			},
			mockFnCache: func(m *MockCache) {
				m.On("Get", "alias").Return("", false)
				m.On("Set", "alias", "http://kimoutofkim.blog", time.Hour)
			},
			want:    "http://kimoutofkim.blog",
			wantErr: false,
		},
		{
			name: "Cache miss, database miss",
			args: args{
				ctx:         context.Background(),
				customAlias: "nonExistentAlias",
			},
			mockFn: func(m *MockQueries) {
				m.On("GetOriginalURL", mock.Anything, "nonExistentAlias").Return("", errors.New("failed to get original url"))
			},
			mockFnCache: func(m *MockCache) {
				m.On("Get", "nonExistentAlias").Return("", false)
			},
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockQueries := &MockQueries{}
			mockCache := &MockCache{}
			tt.mockFn(mockQueries)
			tt.mockFnCache(mockCache)

			s := &URLService{
				queries: mockQueries,
				cfg:     cfg,
				cache:   mockCache,
			}

			got, err := s.GetOriginalURL(tt.args.ctx, tt.args.customAlias)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}

			mockQueries.AssertExpectations(t)
			mockCache.AssertExpectations(t)
		})
	}
}
