package usecases

import (
	"context"
	"testing"
	"time"

	"github.com/dalmazox/accountifier/config"
	authv1 "github.com/dalmazox/accountifier/gen/go/auth/v1"
	"github.com/dalmazox/accountifier/internal/models"
	"github.com/dalmazox/accountifier/internal/repositories/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestLoginShouldAuthenticate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	const (
		email          = "email@test.com"
		rawPassword    = "123456"
		hashedPassword = "$2a$12$Mgn4s5HPwGSAICgteBSpdunxfTNPnMA2LEjqn8j2HVLdJYCikyK4u"
	)

	configMock := &config.Config{App: config.AppConfig{Secret: "fake-secret"}}
	userRepoMock := mocks.NewMockIUserRepository(ctrl)
	userTokenRepoMock := mocks.NewMockIUserTokenRepository(ctrl)
	txMock := mocks.NewMockITx(ctrl)

	user := &models.User{
		UUID:         "test-uuid",
		Email:        email,
		PasswordHash: hashedPassword,
	}

	userRepoMock.EXPECT().BeginTx().Return(txMock, nil).Times(1)
	userRepoMock.EXPECT().GetUserByEmail(gomock.Any(), email, txMock).Return(user, nil).Times(1)
	userTokenRepoMock.EXPECT().UpsertToken(gomock.Any(), gomock.Any(), txMock).Return(nil).Times(1)
	txMock.EXPECT().Commit().Return(nil).Times(1)

	useCase := NewLoginUseCase(configMock, userRepoMock, userTokenRepoMock)
	request := &authv1.LoginRequest{Email: email, Password: rawPassword}

	response, err := useCase.Login(context.Background(), request)

	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Contains(t, response.AccessToken, "ey")
	assert.NotEmpty(t, response.RefreshToken)
	assert.WithinDuration(t, time.Now().Add(1*time.Hour), response.ExpiresAt.AsTime(), 5*time.Second)
}
