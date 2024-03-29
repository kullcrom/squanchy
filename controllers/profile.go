package controllers

import (
	"encoding/json"
	"github.com/kullcrom/squanchy/db"
	"log"
	"net/http"
)

type username struct {
	Email string
}

//ProfileHandler is a custom handler that takes a global Redis pool connection
type ProfileHandler struct{}

//HandleProfile ...
func (h ProfileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodPost:
		var user username
		err := json.NewDecoder(r.Body).Decode(&user)
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
