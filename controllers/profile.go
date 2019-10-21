package controllers

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/gomodule/redigo/redis"
	"github.com/kullcrom/squanchy/db"
	"github.com/kullcrom/squanchy/services"
	"github.com/kullcrom/squanchy/types"
	"log"
	"net/http"
	"os"
)

type username struct {
	Email string
}

//ProfileHandler is a custom handler that takes a global Redis pool connection
type ProfileHandler struct {
	Pool *redis.Pool
}

//HandleProfile ...
func (h ProfileHandler) HandleProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodPost:
		// Authenticate
		cookie, err := r.Cookie("jwt_token")
		if err != nil {
			if err == http.ErrNoCookie {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		jwtKey, exists := os.LookupEnv("JWT_SECRET_KEY")
		if !exists {
			panic("ERROR: JWT_SECRET_KEY not found")
		}

		var jwtSecretKey = []byte(jwtKey)
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

		var user username
		err = json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		userDetails, err := db.GetUserByEmail(user.Email)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(userDetails)
	case http.MethodOptions:
		w.WriteHeader(http.StatusOK)
		return
	default:
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

//HandleProfileUpdate ...
func (h ProfileHandler) HandleProfileUpdate(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var user types.User
		json.NewDecoder(r.Body).Decode(&user)

		authenticated, err := auth.Authenticate(r, h.Pool)
		if err != nil || authenticated != user.Email {
			log.Println(err)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		err = db.UpdateUser(&user)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Profile could not be updated"))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Profile updated successfully"))
		return
	case http.MethodOptions:
		w.WriteHeader(http.StatusOK)
		return
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(r.Method + " HTTP method not accepted"))
		return
	}
}
