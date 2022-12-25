package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	// JSON is used for postman
	ID       string    `json:"id"`
	UID      string    `json:"uid"`
	Title    string    `json:"title"`
	Director *Director `json:""`
}

type Director struct {
	// JSON is used for postman
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Slice to store movies
var movies []Movie

// Handeler Functions

func getMovies() {}

func getMovie() {}

func createMovie() {}

func updateMovie() {}

func deleteMovie() {}

func main() {

	// Lets create some movies beforehand so getMovies work

	movies = append(movies, Movie{ID: "1", UID: "2536113", Title: "First Movie",
		Director: &Director{Firstname: "John", Lastname: "Doe"}})

	movies = append(movies, Movie{ID: "2", UID: "2536123", Title: "Second Movie",
		Director: &Director{Firstname: "Aman", Lastname: "Sharma"}})

	r := mux.NewRouter()

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Starting  server on the port 8080\n")
	log.Fatal(http.ListenAndServe("8080", r))

}
