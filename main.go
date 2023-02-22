package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Creating necessary structs
type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Init movies var as a slice Movie struct
var movies []Movie

func main() {
	//init router
	r := mux.NewRouter()

	// Mock data

	movies = append(movies, Movie{
		ID:       "1",
		Isbn:     "12345",
		Title:    "Movie One",
		Director: &Director{Firstname: "John", Lastname: "Doe"},
	})

	movies = append(movies, Movie{
		ID:       "2",
		Isbn:     "12346",
		Title:    "Movie Two",
		Director: &Director{Firstname: "Steve", Lastname: "Smith"}})

	//creating the necessary routes
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovieById).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	//Starting server
	fmt.Printf("Starting server on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}

// func get all movies
func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)

}

// func delete a movie
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)

}

// func to get a movie by id
func getMovieById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}
