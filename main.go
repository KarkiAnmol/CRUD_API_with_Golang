package main

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

func main() {
	// Your main function is currently empty.
	// You can add your application logic here.
}
