package url

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/nurmuh-alhakim18/url-shortener-api/config"
	"github.com/nurmuh-alhakim18/url-shortener-api/pkg/redis"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func (m *MockQueries) GetOriginalURL(ctx context.Context, customAlias string) (string, error) {
	args := m.Called(ctx, customAlias)
	return args.String(0), args.Error(1)
}

func TestURLService_GetOriginalURL(t *testing.T) {
	cfg := config.Config{}

	type args struct {
		ctx         context.Context
		customAlias string
	}

	tests := []struct {
		name          string
		args          args
		mockFnQueries func(m *MockQueries)
		mockFnRedis   func(m *redis.MockRedisClient)
		want          string
		wantErr       bool
	}{
		{
			name: "Redis hit for original URL",
			args: args{
				ctx:         context.Background(),
				customAlias: "alias",
			},
			mockFnQueries: func(m *MockQueries) {
			},
			mockFnRedis: func(m *redis.MockRedisClient) {
				m.On("Get", mock.Anything, "alias").Return("http://kimoutofkim.blog", nil)
			},
			want:    "http://kimoutofkim.blog",
			wantErr: false,
		},
		{
			name: "Redis miss, database hit for original URL",
			args: args{
				ctx:         context.Background(),
				customAlias: "alias",
			},
			mockFnQueries: func(m *MockQueries) {
				m.On("GetOriginalURL", mock.Anything, "alias").Return("http://kimoutofkim.blog", nil)
			},
			mockFnRedis: func(m *redis.MockRedisClient) {
				m.On("Get", mock.Anything, "alias").Return("", errors.New("redis miss"))
				m.On("Set", mock.Anything, "alias", "http://kimoutofkim.blog", time.Hour).Return(nil)
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
			mockFnQueries: func(m *MockQueries) {
				m.On("GetOriginalURL", mock.Anything, "nonExistentAlias").Return("", errors.New("url not found"))
			},
			mockFnRedis: func(m *redis.MockRedisClient) {
				m.On("Get", mock.Anything, "nonExistentAlias").Return("", errors.New("cache miss"))
			},
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockQueries := &MockQueries{}
			mockRedis := &redis.MockRedisClient{}
			tt.mockFnQueries(mockQueries)
			tt.mockFnRedis(mockRedis)

			s := &URLService{
				queries: mockQueries,
				cfg:     cfg,
				redis:   mockRedis,
			}

			got, err := s.GetOriginalURL(tt.args.ctx, tt.args.customAlias)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}

			mockQueries.AssertExpectations(t)
			mockRedis.AssertExpectations(t)
		})
	}
}
