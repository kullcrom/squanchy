package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/kullcrom/squanchy/types"
	"time"
)

//GenerateToken generates a new JSON Web Token for user login or registration
func GenerateToken(jwtKey string, expiration time.Time, username string) (string, error) {
	var jwtSecretKey = []byte(jwtKey)

	claims := &types.Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiration.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
