package controllers

import (
	"encoding/json"
	"github.com/kullcrom/squanchy/db"
	"github.com/kullcrom/squanchy/services"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"time"
)

type loginCreds struct {
	Username string
	Password string
}

//LoginHandler is a custom handler that takes a JWT secret key
type LoginHandler struct {
	JWTSecretKey string
}

//HandleLogin ...
func (h LoginHandler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	// Check Request Method. Handle only POST for login.
	switch r.Method {
	case http.MethodPost:
		// Query DB for existing User. If found, compare hashed password with provided password.
		// If not found, then return "invalid username/password"
		var creds loginCreds
		err := json.NewDecoder(r.Body).Decode(&creds)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		user, err := db.GetUserByEmail(creds.Username)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid username/password"))
			return
		}

		pwError := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password))
		if pwError != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid username/password"))
			return
		}

		expiration := time.Now().Add(5 * time.Minute)
		tokenString, err := auth.GenerateToken(h.JWTSecretKey, expiration, creds.Username)
		if err != nil || tokenString == "" {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		cookie := http.Cookie{
			Name:    "jwt_token",
			Value:   tokenString,
			Expires: expiration,
		}
		http.SetCookie(w, &cookie)
		w.Write([]byte("jwt_token: " + tokenString))
	case http.MethodOptions:
		w.WriteHeader(http.StatusOK)
		return
	default:
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
