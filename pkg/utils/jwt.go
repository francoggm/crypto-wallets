package utils

import (
	"github/francoggm/crypto-wallets/internal/auth"
	"github/francoggm/crypto-wallets/internal/models"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	UserID   int
	UserRole string
	jwt.RegisteredClaims
}

func GenerateTokenJWT(exp int, secretKey string, user *models.User) (string, error) {
	claims := Claims{
		UserID:   user.ID,
		UserRole: user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{
				Time: time.Now().Add(time.Duration(exp) * time.Hour),
			},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

func ValidateTokenJWT(token, secretKey string) (*Claims, error) {
	claims := Claims{}
	parsedToken, err := jwt.ParseWithClaims(token, &claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	if !parsedToken.Valid {
		return nil, auth.ErrTokenNotValid
	}

	return &claims, nil
}
