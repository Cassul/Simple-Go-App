package api

import (
	"fmt"
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
