package usecases

import (
	"context"
	"errors"

	"github.com/dalmazox/accountifier/config"
	authv1 "github.com/dalmazox/accountifier/gen/go/auth/v1"
	"github.com/dalmazox/accountifier/internal/models"
	"github.com/dalmazox/accountifier/internal/repositories"
	"github.com/dalmazox/accountifier/internal/utils"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const tokenType = "Bearer"

type RefreshTokenUseCase struct {
	config        *config.Config
	userTokenRepo repositories.IUserTokenRepository
}

func NewRefreshTokenUseCase(config *config.Config, userTokenRepo repositories.IUserTokenRepository) *RefreshTokenUseCase {
	return &RefreshTokenUseCase{
		config:        config,
		userTokenRepo: userTokenRepo,
	}
}

func (useCase *RefreshTokenUseCase) RefreshToken(ctx context.Context, request *authv1.RefreshTokenRequest) (*authv1.RefreshTokenResponse, error) {
	invalidRefreshTokenError := errors.New("invalid refresh token")
	tx, err := useCase.userTokenRepo.BeginTx()
	if err != nil {
		return nil, errors.New("failed to open database")
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	request.RefreshToken = utils.HashToken(request.RefreshToken)
	userToken, err := useCase.userTokenRepo.GetUserTokenByRefreshToken(ctx, request.RefreshToken, tx)
	if err != nil {
		return nil, invalidRefreshTokenError
	}

	token, err := utils.GenerateJwt(userToken.UserId, useCase.config)
	if err != nil {
		return nil, errors.New("failed to generate token")
	}

	refreshToken, err := utils.GenerateRefreshToken(userToken.UserId)
	if err != nil {
		return nil, errors.New("failed to generate refresh token")
	}

	newUserToken := &models.UserToken{
		UserId:                userToken.UserId,
		Token:                 token.Token,
		RefreshToken:          refreshToken.HashedToken,
		TokenExpiresAt:        token.ExpiresAt,
		RefreshTokenExpiresAt: refreshToken.ExpiresAt,
	}

	err = useCase.userTokenRepo.UpsertToken(ctx, *newUserToken, tx)
	if err != nil {
		return nil, errors.New("failed to upsert token")
	}

	err = tx.Commit()
	if err != nil {
		return nil, errors.New("failed to commit transaction")
	}

	return &authv1.RefreshTokenResponse{
		AccessToken:  token.Token,
		RefreshToken: refreshToken.OriginalToken,
		TokenType:    tokenType,
		ExpiresAt:    timestamppb.New(token.ExpiresAt),
	}, nil
}
