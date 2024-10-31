package url

import (
	"context"
	"errors"
	"testing"

	"github.com/nurmuh-alhakim18/url-shortener-api/config"
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
		name    string
		args    args
		mockFn  func(m *MockQueries)
		want    string
		wantErr bool
	}{
		{
			name: "Successful getting original URL",
			args: args{
				ctx:         context.Background(),
				customAlias: "alias",
			},
			mockFn: func(m *MockQueries) {
				m.On("GetOriginalURL", mock.Anything, "alias").Return("http://kimoutofkim.blog", nil)
			},
			want:    "http://kimoutofkim.blog",
			wantErr: false,
		},
		{
			name: "Custom alias not found",
			args: args{
				ctx:         context.Background(),
				customAlias: "nonExistentAlias",
			},
			mockFn: func(m *MockQueries) {
				m.On("GetOriginalURL", mock.Anything, "nonExistentAlias").Return("", errors.New("failed to get orinal url"))
			},
			want:    "",
			wantErr: true,
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

			got, err := s.GetOriginalURL(tt.args.ctx, tt.args.customAlias)
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
