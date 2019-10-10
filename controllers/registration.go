package controllers

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"github.com/kullcrom/squanchy/db"
	"github.com/kullcrom/squanchy/services"
	"log"
	"net/http"
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

func (h RegistrationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

		pool := h.Pool
		conn := pool.Get()
		defer conn.Close()

		sessionToken, err := auth.GenerateSessionToken()
		if err != nil || sessionToken == "" {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("ERROR: Session token cannot be nil"))
			return
		}

		_, err = conn.Do("SETEX", sessionToken, "300", user.Username)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		cookie := http.Cookie{
			Name:    "session_token",
			Value:   sessionToken,
			Expires: time.Now().Add(5 * time.Minute),
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
