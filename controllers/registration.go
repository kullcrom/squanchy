package controllers

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/gomodule/redigo/redis"
	"github.com/kullcrom/squanchy/db"
	"github.com/kullcrom/squanchy/types"
	"log"
	"net/http"
	"os"
	"time"
)

type userRegistration struct {
	Username        string
	Password        string
	ConfirmPassword string
}

//RegistrationHandler is a custom handler that takes a global Redis pool connection
type RegistrationHandler struct {
	Pool *redis.Pool
}

//HandleRegistration ...
func (h RegistrationHandler) HandleRegistration(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodPost:
		var user userRegistration
		json.NewDecoder(r.Body).Decode(&user)

		if user.Password != user.ConfirmPassword {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Passwords do not match"))
			return
		}

		err := db.CreateUser(user.Username, user.Password)
		if err != nil {
			log.Println(err)
			if err.Error() == "ERROR: User already exists" {
				w.WriteHeader(http.StatusBadRequest)
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
			w.Write([]byte(err.Error()))
			return
		}

		jwtKey, exists := os.LookupEnv("JWT_SECRET_KEY")
		if !exists {
			panic("ERROR: JWT_SECRET_KEY not found")
		}

		var jwtSecretKey = []byte(jwtKey)

		expiration := time.Now().Add(5 * time.Minute)
		claims := &types.Claims{
			Username: user.Username,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expiration.Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(jwtSecretKey)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		cookie := http.Cookie{
			Name:    "jwt_token",
			Value:   tokenString,
			Expires: expiration,
		}

		http.SetCookie(w, &cookie)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("User created successfully"))
		return
	case http.MethodOptions:
		w.WriteHeader(http.StatusOK)
		return
	default:
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
