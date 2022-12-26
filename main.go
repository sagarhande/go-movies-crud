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

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
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
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000)) // Generate random number

	movies = append(movies, movie)

	// return that newly created movie
	json.NewEncoder(w).Encode(movie)

}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for i, item := range movies {
		if item.ID == params["id"] {
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movies[i].UID = movie.UID
			movies[i].Title = movie.Title
			movies[i].Director = movie.Director
			json.NewEncoder(w).Encode(movies[i])
			return
		}
	}
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get all params/values inside request by using below
	params := mux.Vars(r)

	// for loop to check id
	for i, item := range movies {
		if item.ID == params["id"] {
			/*Without the dots you are trying to add a byte slice as an element of a byte slice. This is not allowed. By adding dots you essentially unpack the second argument into a bunch of individual bytes and add them one at a time.*/
			movies = append(movies[:i], movies[i+1:]...)
			break

		}
	}

	// Return all movies
	json.NewEncoder(w).Encode(movies)
}

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
	log.Fatal(http.ListenAndServe(":8080", r))

}
