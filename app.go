package main

import (
	"net/http"

	"github.com/cassul/api"
	"github.com/cassul/bigtable"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/about", api.AboutHandler)
	r.HandleFunc("/", api.MainPageHandler)
	http.Handle("/", r)
	bigtable.EmulateBigTable()

	http.ListenAndServe(":8080", r)
}
