package usecases

import (
	"context"
	"errors"
	"time"

	authv1 "github.com/dalmazox/accountifier/gen/go/auth/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type LoginUseCase struct{}

func NewLoginUseCase() *LoginUseCase {
	return &LoginUseCase{}
}

func (useCase *LoginUseCase) Login(ctx context.Context, request *authv1.LoginRequest) (*authv1.LoginResponse, error) {
	if request.Email == "test@qa.com" && request.Password == "123456" {
		return &authv1.LoginResponse{
			AccessToken:  "access_token_example",
			RefreshToken: "refresh_token_example",
			TokenType:    "Bearer",
			ExpiresAt:    timestamppb.New(time.Now().Add(1 * time.Hour)),
		}, nil
	}

	return nil, errors.New("invalid credentials")
}
