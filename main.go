package main

import (
	"encoding/json"
	"fmt"
	"github.com/kullcrom/squanchy/attomdata/basicprofile"
	"github.com/kullcrom/squanchy/db"
	"github.com/kullcrom/squanchy/types"
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
	user1 := new(types.User)
	user1.Email = "name@email.com"
	user1.FirstName = "Billy-Bob"
	user1.LastName = "Thorton"

	// db.CreateUser(*user1)
	// db.DeleteUser(*user1)
	user2 := db.GetUserByID(2)
	output, err := json.Marshal(user2)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(output))

	http.HandleFunc("/", basicprofile.GetBasicProfile)
	http.HandleFunc("/address", address)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
