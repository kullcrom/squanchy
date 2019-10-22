package controllers

import (
	"encoding/json"
	"github.com/kullcrom/squanchy/db"
	"github.com/kullcrom/squanchy/types"
	"net/http"
)

//UpdateProfileHandler is a custom handler for profile updates
type UpdateProfileHandler struct{}

//HandleProfileUpdate ...
func (h UpdateProfileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodPost:
		var user types.User
		json.NewDecoder(r.Body).Decode(&user)

		err := db.UpdateUser(&user)
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
