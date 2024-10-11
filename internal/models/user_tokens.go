package models

import "time"

type UserToken struct {
	UUID                  string    `db:"id"`
	UserId                string    `db:"user_id"`
	Token                 string    `db:"token"`
	RefreshToken          string    `db:"refresh_token"`
	CreatedAt             time.Time `db:"created_at"`
	UpdatedAt             time.Time `db:"updated_at"`
	TokenExpiresAt        time.Time `db:"token_expires_at"`
	RefreshTokenExpiresAt time.Time `db:"refresh_token_expires_at"`
	Revoked               bool      `db:"revoked"`
}
