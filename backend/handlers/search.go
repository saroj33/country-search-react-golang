package handlers

import (
	"country-search-backend/searcher"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// SearchHandler take cares of all the search operations related to countries
func SearchHandler(w http.ResponseWriter, r *http.Request) {
	// Setting CORS. This can be done on a middleware. Ans Security shoiuld be added if required
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	// Getting text query string from request url
	searchText := r.FormValue("text")
	// Loading our file that contains all the data
	file, _ := ioutil.ReadFile("datasource/countries.json")
	var fullList searcher.CountryList
	if err := json.Unmarshal([]byte(file), &fullList); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
		return
	}

	// Calling Search method and passing the searchText
	res, err := fullList.Search(searchText)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
	return

}
