package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// the Movie struct that is going to be used throughout the API
type Movie struct {
	MovieTitle      string   `json:"title"`
	Year            int      `json:"release_year"`
	MainActors      []string `json:"actors"`
	MoviePoster     string   `json:"posters"`
	SimilarMovieIDs []string `json:"similar_ids"`
}

// easy space removal
func removeSpaces(str string) string {
	return strings.TrimSpace(str)
}

// converting struct Movie to JSON
func createJson(mov Movie) string {
	if mov.MovieTitle != "" {
		json, error := json.Marshal(mov)
		if error != nil {
			fmt.Println("Error: Converting struct to JSON")
		}
		return string(json)
	}
	return ""
}

// retrieving Movie info from the website from the respective tags
func retrieveInfo(movieUrl string) Movie {
	var movie Movie
	var year int
	//var err error
	var actorlist []string
	var similarlist []string
	movie = Movie{}

	webpage, err := goquery.NewDocument(movieUrl)
	if err != nil {
		fmt.Println("Error: Loading the webpage")
	}

	// fetching similar movie ids
	webpage.Find("div").Find(".dv-universal-hover-enabled").Each(func(index int, w *goquery.Selection) {
		similarID, _ := w.Attr("data-asin")
		similarID = removeSpaces(similarID)
		similarlist = append(similarlist, removeSpaces(similarID))
	})
	movie.SimilarMovieIDs = similarlist

	// check if the movie details are fetched correctly
	if similarlist == nil {
		return Movie{}
	}

	// fetching movie title
	movie.MovieTitle = removeSpaces(webpage.Find("#aiv-content-title").Before("span").Text())

	// fetching main movie actors
	actorlist = strings.Split(removeSpaces(webpage.Find("dd").Before("dt").First().Text()), ",")
	movie.MainActors = actorlist

	// fetching movie release year
	year, err = strconv.Atoi(removeSpaces(webpage.Find(".release-year").Before("span").Text()))
	if err != nil {
		fmt.Println("Error: Cannot retrieve movie release year")
	}

	movie.Year = year

	// fetching movie poster
	movie.MoviePoster, _ = webpage.Find(".dp-img-bracket").Find(".dp-meta-icon-container").Find("img").Attr("src")

	return movie
}
