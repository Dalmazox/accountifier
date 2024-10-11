package models

import "time"

type User struct {
	UUID           string     `db:"id"`
	Email          string     `db:"email"`
	Username       string     `db:"username"`
	PasswordHash   string     `db:"password_hash"`
	CreatedAt      time.Time  `db:"created_at"`
	UpdatedAt      time.Time  `db:"updated_at"`
	Active         bool       `db:"is_active"`
	LastLogin      *time.Time `db:"last_login"`
	FailedAttempts *int       `db:"failed_attempts"`
	LockedUntil    *time.Time `db:"locked_until"`
}
