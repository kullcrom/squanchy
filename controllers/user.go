package controllers

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/gomodule/redigo/redis"
	"github.com/kullcrom/squanchy/db"
	"github.com/kullcrom/squanchy/types"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type loginCreds struct {
	Username string
	Password string
}

//LoginHandler is a custom handler that takes a global Redis Pool connection
type LoginHandler struct {
	Pool *redis.Pool
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

		jwtKey, exists := os.LookupEnv("JWT_SECRET_KEY")
		if !exists {
			panic("ERROR: JWT_SECRET_KEY not found")
		}

		var jwtSecretKey = []byte(jwtKey)

		expiration := time.Now().Add(5 * time.Minute)
		claims := &types.Claims{
			Username: creds.Username,
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
		w.Write([]byte("jwt_token: " + tokenString))
	case http.MethodOptions:
		w.WriteHeader(http.StatusOK)
		return
	default:
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

//GetUsers ...
func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := []types.User{}

	users, err := db.GetAllUsers()
	if err != nil {
		log.Fatal(err)
	}

	output, err := json.Marshal(users)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(output)
}

//GetUserByID ...
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	idString := strings.TrimPrefix(r.URL.Path, "/users/")

	id, err := strconv.Atoi(idString)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("Invalid ID. Must be of type integer.")
	} else {
		user, err := db.GetUserByID(id)
		if err != nil {
			log.Println(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		output, err := json.Marshal(user)
		if err != nil {
			log.Println(err)
		}
		w.Write(output)
	}
}
