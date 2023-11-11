package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Movie represents a movie with its details.
type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

// Director represents the director of a movie.
type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}
func deleteMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applicaiton/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}
func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}
func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(movie)
	movie.ID = strconv.Itoa(rand.Intn(10000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}
func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
		}
	}
}

func main() {
	// Create a new Gorilla mux router
	r := mux.NewRouter()
	movies = append(movies, Movie{
		ID:    "1",
		Isbn:  "1234",
		Title: "movie one",
		Director: &Director{
			FirstName: "Director",
			LastName:  "One",
		},
	})
	movies = append(movies, Movie{
		ID:    "2",
		Isbn:  "2345",
		Title: "movie two",
		Director: &Director{
			FirstName: "Director",
			LastName:  "Two",
		},
	})
	// Define routes and associated handler functions
	// Create a new movie
	r.HandleFunc("/movies", createMovie).Methods("POST")
	// Get all movies
	r.HandleFunc("/movies", getMovies).Methods("GET")
	// Get a specific movie by ID
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	// Update a movie by ID
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	// Delete a movie by ID
	r.HandleFunc("/movies/{id}", deleteMovies).Methods("DELETE")

	fmt.Println("Starting Server at port 8084")
	log.Fatal(http.ListenAndServe(":8084", r))

}
