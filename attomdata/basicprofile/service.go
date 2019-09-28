package basicprofile

import (
	"encoding/json"
	"fmt"
	"github.com/motemen/go-loghttp"
	"io/ioutil"
	"log"
	"net/http"
)

//GetBasicProfile ...
func GetBasicProfile(w http.ResponseWriter, r *http.Request) {
	apiKey := "6ce02f2add01852402f0c48a1bea1af5"
	address1 := r.URL.Query().Get("address1")
	address2 := r.URL.Query().Get("address2")

	client := &http.Client{
		Transport: &loghttp.Transport{},
	}
	req, err := http.NewRequest("GET", "https://api.gateway.attomdata.com/propertyapi/v1.0.0/property/basicprofile", nil)
	query := req.URL.Query()
	query.Add("address1", address1)
	query.Add("address2", address2)
	req.URL.RawQuery = query.Encode()

	req.Header["apikey"] = []string{apiKey}
	req.Header.Add("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	results, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	response := BasicProfile{}
	json.Unmarshal(results, &response)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(response.Status.Msg)
	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(req.URL.String())
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
