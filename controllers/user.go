package controllers

import (
	"encoding/json"
	"github.com/kullcrom/squanchy/db"
	"github.com/kullcrom/squanchy/types"
	"log"
	"net/http"
	"strconv"
	"strings"
)

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
