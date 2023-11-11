package main

import "github.com/gorilla/mux"

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

var Movies []Movie

func main() {
	// Create a new Gorilla mux router
	r := mux.NewRouter()

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

}
