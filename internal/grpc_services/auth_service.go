package grpcservices

import (
	"context"

	authv1 "github.com/dalmazox/accountifier/gen/go/auth/v1"
	"github.com/dalmazox/accountifier/internal/usecases"
)

type AuthServiceServer struct {
	authv1.UnimplementedAuthServiceServer
	loginUseCase *usecases.LoginUseCase
}

func NewAuthServiceServer(loginUseCase *usecases.LoginUseCase) *AuthServiceServer {
	return &AuthServiceServer{
		loginUseCase: loginUseCase,
	}
}

func (server *AuthServiceServer) Login(ctx context.Context, request *authv1.LoginRequest) (*authv1.LoginResponse, error) {
	return server.loginUseCase.Login(ctx, request)
}
