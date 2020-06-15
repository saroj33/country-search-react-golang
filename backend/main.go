package main

import (
	"fmt"
	"log"
	"net/http"

	"country-search-backend/handlers"

	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Nothing here. Use appropirate service url")
}
func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", home)
	// This is our single useful API endpoint
	router.HandleFunc("/search", handlers.SearchHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":2020", router))

}
