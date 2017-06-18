package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	// router's unique handler functions to cover different scenarios of API's usage
	router.HandleFunc("/", Index)
	router.HandleFunc("/movie/amazon/", Index)
	router.HandleFunc("/movie/amazon/{countryAbbr}/{movieId}", movieHandler)

	log.Fatal(http.ListenAndServe(":8080", router))
}

// handles access to localhost:8080 and gives directives to the user
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Please enter a region (com or co.uk or de) after /movie/amazon/\n")
	fmt.Fprintf(w, "Please enter an Amazon Movie ID after /movie/amazon/{countryAbbrevation}/")
}

// supports different Amazon websites like .com .co.uk .de.
func movieHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	movieId := vars["movieId"]
	countryAbbr := vars["countryAbbr"]

	// url is formatted with the Amazon's region and movie's id
	// it is then converted into JSON with a function call to createJson that is in retrieve_data.go
	url := "http://www.amazon." + countryAbbr + "/gp/product/" + string(movieId)
	movie := retrieveInfo(url)
	movieJson := createJson(movie)

	// providing the JSON data to the user if no problem has occured.
	if movieJson != "" {
		fmt.Fprintln(w, movieJson)
		return
	}

	fmt.Fprintln(w, "the movieId you've in the link was either invalid for this region or it does not exist.")
	return
}
