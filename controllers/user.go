package controllers

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"github.com/kullcrom/squanchy/db"
	"github.com/kullcrom/squanchy/services"
	"github.com/kullcrom/squanchy/types"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
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

//UserLogin ...
func (h LoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	// Check Request Method. Handle only POST for login.
	switch r.Method {
	case http.MethodPost:
		pool := h.Pool
		conn := pool.Get()
		defer conn.Close()

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

		// If creds are valid, generate a token and create session to log user in. Save session info
		// to Redis for quick session authentication and set-cookie for session.
		sessionToken, err := auth.GenerateSessionToken()
		if err != nil || sessionToken == "" {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error: Session token cannot be nil"))
			return
		}

		_, err = conn.Do("SETEX", sessionToken, "300", creds.Username)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		cookie := http.Cookie{
			Name:    "session_token",
			Value:   sessionToken,
			Expires: time.Now().Add(5 * time.Minute),
		}
		http.SetCookie(w, &cookie)
		w.Write([]byte("Token: " + sessionToken))
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
