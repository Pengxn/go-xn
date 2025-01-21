package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/Pengxn/go-xn/src/config"
)

type Claims struct {
	UID      int    `json:"uid"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func NewClaims(uid int, username string) Claims {
	now := time.Now()

	claims := Claims{
		UID:      uid,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt: jwt.NewNumericDate(now),
		},
	}
	// Set expiration time if config value is set, the value is in days.
	if config.Config.Server.JwtExp > 0 {
		claims.ExpiresAt = jwt.NewNumericDate(now.AddDate(0, 0, config.Config.Server.JwtExp))
	}

	return claims
}

func TokenFromClaims(claims Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.Config.Server.JwtToken))
}
