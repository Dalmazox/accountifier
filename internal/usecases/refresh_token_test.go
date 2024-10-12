package usecases

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/dalmazox/accountifier/config"
	authv1 "github.com/dalmazox/accountifier/gen/go/auth/v1"
	"github.com/dalmazox/accountifier/internal/models"
	"github.com/dalmazox/accountifier/internal/repositories/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestRefreshToken(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	configMock := &config.Config{App: config.AppConfig{Secret: "fake-secret"}}
	userTokenRepoMock := mocks.NewMockIUserTokenRepository(ctrl)
	txMock := mocks.NewMockITx(ctrl)

	const (
		refreshToken       = "45ed12eb583e2288c74136887ed51eab004986f55e236b601d3ab81ad394cf40"
		hashedRefreshToken = "a655f99a9cfaeda410eebd944494f6d53803127ab2ab2b922035f9ff771ac46f"
	)

	tests := []struct {
		name  string
		args  *authv1.RefreshTokenRequest
		setup func()
		err   error
	}{
		{
			name: "should return new access token",
			args: &authv1.RefreshTokenRequest{RefreshToken: refreshToken},
			setup: func() {
				userTokenRepoMock.EXPECT().BeginTx().Return(txMock, nil).Times(1)
				userTokenRepoMock.EXPECT().GetUserTokenByRefreshToken(gomock.Any(), hashedRefreshToken, txMock).Return(&models.UserToken{
					UUID:         "test-uuid",
					UserId:       "test-user-id",
					RefreshToken: hashedRefreshToken,
				}, nil).Times(1)
				userTokenRepoMock.EXPECT().UpsertToken(gomock.Any(), gomock.Any(), txMock).Return(nil).Times(1)
				txMock.EXPECT().Commit().Return(nil).Times(1)
			},
			err: nil,
		},
		{
			name: "should return error when refresh token not found",
			args: &authv1.RefreshTokenRequest{RefreshToken: refreshToken},
			setup: func() {
				userTokenRepoMock.EXPECT().BeginTx().Return(txMock, nil).Times(1)
				userTokenRepoMock.EXPECT().GetUserTokenByRefreshToken(gomock.Any(), gomock.Any(), txMock).Return(nil, sql.ErrNoRows)
				txMock.EXPECT().Rollback().Return(nil).Times(1)
			},
			err: errors.New("invalid refresh token"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()

			useCase := NewRefreshTokenUseCase(configMock, userTokenRepoMock)
			response, err := useCase.RefreshToken(context.Background(), tt.args)

			if tt.err != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.err, err)
				return
			}

			assert.NotNil(t, response)
		})
	}

}
