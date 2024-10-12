package repositories

import (
	"context"

	"github.com/dalmazox/accountifier/internal/models"
	"github.com/jmoiron/sqlx"
)

type IUserRepository interface {
	GetUserByEmail(ctx context.Context, email string, tx ITx) (*models.User, error)
	BeginTx() (ITx, error)
}

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (repo *UserRepository) GetUserByEmail(ctx context.Context, email string, tx ITx) (*models.User, error) {
	query := LoadQuery("get_user_by_email")
	var user models.User
	if err := tx.GetContext(ctx, &user, query, email); err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *UserRepository) BeginTx() (ITx, error) {
	return repo.db.Beginx()
}
