package usecases

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/dalmazox/accountifier/config"
	authv1 "github.com/dalmazox/accountifier/gen/go/auth/v1"
	"github.com/dalmazox/accountifier/internal/models"
	"github.com/dalmazox/accountifier/internal/repositories/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	configMock := &config.Config{App: config.AppConfig{Secret: "fake-secret"}}
	userRepoMock := mocks.NewMockIUserRepository(ctrl)
	userTokenRepoMock := mocks.NewMockIUserTokenRepository(ctrl)
	txMock := mocks.NewMockITx(ctrl)

	const (
		email          = "email@test.com"
		rawPassword    = "123456"
		hashedPassword = "$2a$12$Mgn4s5HPwGSAICgteBSpdunxfTNPnMA2LEjqn8j2HVLdJYCikyK4u"
	)

	tests := []struct {
		name  string
		args  *authv1.LoginRequest
		setup func()
		err   error
	}{
		{
			name: "should login",
			args: &authv1.LoginRequest{Email: email, Password: rawPassword},
			setup: func() {
				userRepoMock.EXPECT().BeginTx().Return(txMock, nil).Times(1)
				userRepoMock.EXPECT().GetUserByEmail(gomock.Any(), email, txMock).Return(&models.User{
					UUID:         "test-uuid",
					Email:        email,
					PasswordHash: hashedPassword,
				}, nil).Times(1)
				userTokenRepoMock.EXPECT().UpsertToken(gomock.Any(), gomock.Any(), txMock).Return(nil).Times(1)
				txMock.EXPECT().Commit().Return(nil).Times(1)
			},
			err: nil,
		},
		{
			name: "should return invalid credentials error",
			args: &authv1.LoginRequest{Email: "wrong-email", Password: rawPassword},
			setup: func() {
				userRepoMock.EXPECT().BeginTx().Return(txMock, nil).Times(1)
				userRepoMock.EXPECT().GetUserByEmail(gomock.Any(), "wrong-email", txMock).Return(nil, errors.New("")).Times(1)
				txMock.EXPECT().Rollback().Return(nil).Times(1)
			},
			err: errors.New("invalid credentials"),
		},
		{
			name: "should return invalid credentials error when password is incorrect",
			args: &authv1.LoginRequest{Email: email, Password: "wrong-password"},
			setup: func() {
				userRepoMock.EXPECT().BeginTx().Return(txMock, nil).Times(1)
				userRepoMock.EXPECT().GetUserByEmail(gomock.Any(), email, txMock).Return(&models.User{
					UUID:         "test-uuid",
					Email:        email,
					PasswordHash: hashedPassword,
				}, nil).Times(1)
				txMock.EXPECT().Rollback().Return(nil).Times(1)
			},
			err: errors.New("invalid credentials"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()

			useCase := NewLoginUseCase(configMock, userRepoMock, userTokenRepoMock)
			response, err := useCase.Login(context.Background(), tt.args)

			if tt.err != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.err, err)
				return
			}

			assert.NotNil(t, response)
			assert.Equal(t, response.TokenType, "Bearer")
			assert.Contains(t, response.AccessToken, "ey")
			assert.NotEmpty(t, response.RefreshToken)
			assert.WithinDuration(t, time.Now().Add(1*time.Hour), response.ExpiresAt.AsTime(), 5*time.Second)
		})
	}
}
