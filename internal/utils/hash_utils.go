package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"time"

	"github.com/dalmazox/accountifier/config"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserUUID string
	jwt.RegisteredClaims
}

type TokenJwt struct {
	Token     string
	ExpiresAt time.Time
}

type RefreshToken struct {
	HashedToken   string
	OriginalToken string
	UserUUID      string
	ExpiresAt     time.Time
}

func GenerateJwt(userUUID string, cfg *config.Config) (*TokenJwt, error) {
	expirationTime := time.Now().Add(1 * time.Hour)

	claims := &Claims{
		UserUUID: userUUID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(cfg.App.Secret))

	if err != nil {
		return nil, err
	}

	return &TokenJwt{Token: tokenString, ExpiresAt: expirationTime}, nil
}

func GenerateRefreshToken(userUUID string) (*RefreshToken, error) {
	rawToken := make([]byte, 32)
	_, err := rand.Read(rawToken)
	if err != nil {
		return nil, err
	}

	token := hex.EncodeToString(rawToken)
	expiresAt := time.Now().Add(7 * 24 * time.Hour)

	return &RefreshToken{
		HashedToken:   HashToken(token),
		OriginalToken: token,
		UserUUID:      userUUID,
		ExpiresAt:     expiresAt,
	}, nil
}

func RefreshTokenIsValid(providedToken, storedToken string) bool {
	return HashToken(providedToken) == storedToken
}

func HashToken(token string) string {
	hash := sha256.Sum256([]byte(token))
	return hex.EncodeToString(hash[:])
}
