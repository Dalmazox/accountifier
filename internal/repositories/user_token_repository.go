package repositories

import (
	"context"

	"github.com/dalmazox/accountifier/internal/models"
	"github.com/jmoiron/sqlx"
)

type IUserTokenRepository interface {
	UpsertToken(ctx context.Context, userToken models.UserToken, tx ITx) error
	BeginTx() (ITx, error)
}

type UserTokenRepository struct {
	db *sqlx.DB
}

func NewUserTokenRepository(db *sqlx.DB) *UserTokenRepository {
	return &UserTokenRepository{db: db}
}

func (repo *UserTokenRepository) UpsertToken(ctx context.Context, userToken models.UserToken, tx ITx) error {
	query := LoadQuery("insert_user_token")

	_, err := tx.ExecContext(
		ctx,
		query,
		userToken.UserId,
		userToken.Token,
		userToken.RefreshToken,
		userToken.TokenExpiresAt,
		userToken.RefreshTokenExpiresAt)

	return err
}

func (repo *UserTokenRepository) BeginTx() (ITx, error) {
	return repo.db.Beginx()
}
