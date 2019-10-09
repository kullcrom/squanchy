package controllers

import (
	"encoding/json"
	"log"
	"net/http"
)

//AddressForm ...
type AddressForm struct {
	Street string
	City   string
	State  string
}

func address(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/address" {
		http.Error(w, "404 Not Found.", http.StatusNotFound)
		return
	}

	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	var addressForm AddressForm
	err := json.NewDecoder(r.Body).Decode(&addressForm)
	if err != nil {
		log.Println(err)
	}
	out, _ := json.Marshal(addressForm)
	log.Println(string(out))
}
