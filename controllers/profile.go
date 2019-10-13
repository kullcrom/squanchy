package controllers

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"github.com/kullcrom/squanchy/db"
	"github.com/kullcrom/squanchy/services"
	"github.com/kullcrom/squanchy/types"
	"log"
	"net/http"
)

//ProfileHandler is a custom handler that takes a global Redis pool connection
type ProfileHandler struct {
	Pool *redis.Pool
}

//HandleProfile ...
func (h ProfileHandler) HandleProfile(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		w.WriteHeader(http.StatusOK)
		return
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
