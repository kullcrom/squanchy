package main

import (
	"encoding/json"
	"fmt"
	"github.com/kullcrom/squanchy/controllers"
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
	fmt.Println(string(out))
}

func main() {
	http.HandleFunc("/users", controllers.GetUsers)
	http.HandleFunc("/users/", controllers.GetUserByID)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
