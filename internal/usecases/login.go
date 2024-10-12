package usecases

import (
	"context"
	"errors"

	"github.com/dalmazox/accountifier/config"
	authv1 "github.com/dalmazox/accountifier/gen/go/auth/v1"
	"github.com/dalmazox/accountifier/internal/models"
	"github.com/dalmazox/accountifier/internal/repositories"
	"github.com/dalmazox/accountifier/internal/utils"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type LoginUseCase struct {
	config        *config.Config
	userRepo      repositories.IUserRepository
	userTokenRepo repositories.IUserTokenRepository
}

func NewLoginUseCase(
	config *config.Config,
	userRepo repositories.IUserRepository,
	userTokenRepo repositories.IUserTokenRepository) *LoginUseCase {
	return &LoginUseCase{
		config:        config,
		userRepo:      userRepo,
		userTokenRepo: userTokenRepo,
	}
}

func (useCase *LoginUseCase) Login(ctx context.Context, request *authv1.LoginRequest) (*authv1.LoginResponse, error) {
	const tokenType = "Bearer"
	invalidCredentialsError := errors.New("invalid credentials")

	tx, err := useCase.userRepo.BeginTx()
	if err != nil {
		return nil, errors.New("failed to open database")
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	user, err := useCase.userRepo.GetUserByEmail(ctx, request.Email, tx)
	if err != nil {
		return nil, invalidCredentialsError
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(request.Password))
	if err != nil {
		return nil, invalidCredentialsError
	}

	token, err := utils.GenerateJwt(user.UUID, useCase.config)
	if err != nil {
		return nil, errors.New("could not generate token")
	}

	refreshToken, err := utils.GenerateRefreshToken(user.UUID)
	if err != nil {
		return nil, errors.New("could not generate refresh token")
	}

	userToken := models.UserToken{
		Token:                 token.Token,
		RefreshToken:          refreshToken.HashedToken,
		TokenExpiresAt:        token.ExpiresAt,
		RefreshTokenExpiresAt: refreshToken.ExpiresAt,
		UserId:                user.UUID,
	}
	err = useCase.userTokenRepo.UpsertToken(ctx, userToken, tx)
	if err != nil {
		return nil, errors.New("could not upsert token")
	}

	err = tx.Commit()
	if err != nil {
		return nil, errors.New("could not commit transaction")
	}

	return &authv1.LoginResponse{
		AccessToken:  token.Token,
		RefreshToken: refreshToken.OriginalToken,
		ExpiresAt:    timestamppb.New(token.ExpiresAt),
		TokenType:    tokenType,
	}, nil
}
