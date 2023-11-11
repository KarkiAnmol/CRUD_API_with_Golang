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

var movies []Movie // Slice to store movies

// getMovies handles the GET request for retrieving all movies.
func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

// deleteMovies handles the DELETE request for deleting a movie by ID.
func deleteMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			// Delete the movie from the slice
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

// getMovie handles the GET request for retrieving a movie by ID.
func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			// Encode and send the found movie
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

// createMovie handles the POST request for creating a new movie.
func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	// Generate a unique ID for the new movie
	movie.ID = strconv.Itoa(rand.Intn(10000000))
	// Append the new movie to the slice
	movies = append(movies, movie)
	// Encode and send the created movie
	json.NewEncoder(w).Encode(movie)
}

// updateMovie handles the PUT request for updating a movie by ID.
func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			// Delete the existing movie
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			// Decode the updated movie from the request body
			_ = json.NewDecoder(r.Body).Decode(&movie)
			// Set the ID to the existing ID
			movie.ID = params["id"]
			// Append the updated movie to the slice
			movies = append(movies, movie)
			// Encode and send the updated movie
			json.NewEncoder(w).Encode(movie)
		}
	}
}

func main() {
	// Create a new Gorilla mux router
	r := mux.NewRouter()

	// Pre-populate the movies slice with some initial data
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

	// Define routes and associate them with their corresponding handler functions
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

	// Start the server on port 8084
	fmt.Println("Starting Server at port 8084")
	log.Fatal(http.ListenAndServe(":8084", r))
}
