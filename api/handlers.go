package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Oy!")
	fmt.Printf("It worked")
}

func MainPageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome")
	fmt.Printf("First page is open")
}

func SaveCredentials(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}
	defer r.Body.Close()
	var cred credentials
	err = json.Unmarshal(bodyBytes, &cred)
	if err != nil {
		panic(err)
	}
	fmt.Printf("response - %v", cred)
}
