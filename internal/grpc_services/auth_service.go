package grpcservices

import (
	"context"

	authv1 "github.com/dalmazox/accountifier/gen/go/auth/v1"
	"github.com/dalmazox/accountifier/internal/usecases"
)

type AuthServiceServer struct {
	authv1.UnimplementedAuthServiceServer
	loginUseCase        *usecases.LoginUseCase
	refreshTokenUseCase *usecases.RefreshTokenUseCase
}

func NewAuthServiceServer(loginUseCase *usecases.LoginUseCase, refreshTokenUseCase *usecases.RefreshTokenUseCase) *AuthServiceServer {
	return &AuthServiceServer{
		loginUseCase:        loginUseCase,
		refreshTokenUseCase: refreshTokenUseCase,
	}
}

func (server *AuthServiceServer) Login(ctx context.Context, request *authv1.LoginRequest) (*authv1.LoginResponse, error) {
	return server.loginUseCase.Login(ctx, request)
}

func (server *AuthServiceServer) RefreshToken(ctx context.Context, request *authv1.RefreshTokenRequest) (*authv1.RefreshTokenResponse, error) {
	return server.refreshTokenUseCase.RefreshToken(ctx, request)
}
