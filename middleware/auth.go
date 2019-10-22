package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/kullcrom/squanchy/types"
	"net/http"
)

//AuthMiddleware ...
type AuthMiddleware struct {
	JWTSecretKey string
}

//Authenticate ...
func (authHandler *AuthMiddleware) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
		case http.MethodPost:
			cookie, err := r.Cookie("jwt_token")
			if err != nil {
				if err == http.ErrNoCookie {
					w.WriteHeader(http.StatusUnauthorized)
					return
				}
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			var jwtSecretKey = []byte(authHandler.JWTSecretKey)
			tokenString := cookie.Value
			claims := &types.Claims{}

			token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
				return jwtSecretKey, nil
			})
			if err != nil {
				if err == jwt.ErrSignatureInvalid {
					w.WriteHeader(http.StatusUnauthorized)
					return
				}
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			if !token.Valid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		default:
			next.ServeHTTP(w, r)
		}
	})
}
