package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

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

var movies []Movie

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{
		ID:    "1",
		Isbn:  "344565",
		Title: "Pacific Rim",
		Director: &Director{
			Firstname: "James",
			Lastname:  "Cameron",
		},
	})

	movies = append(movies, Movie{
		ID:    "2",
		Isbn:  "3445456",
		Title: "Dark knight",
		Director: &Director{
			Firstname: "Christopher",
			Lastname:  "Nolan",
		},
	})

	r.HandleFunc("/movies", getMoviesHandler).Methods("GET")
	r.HandleFunc("/movies/", createMovieHandler).Methods("POST")
	r.HandleFunc("/movies/{id}", getByIdHandler).Methods("GET")
	r.HandleFunc("/movies/{id}", updateByIdHandler).Methods("POST")
	r.HandleFunc("/movies/{id}", deleteByIdHandler).Methods("DELETE")

	fmt.Println("Starting movie service at port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func getMoviesHandler(w http.ResponseWriter, r *http.Request) {
	log.Fatal("Not implemented")
	return
}

func createMovieHandler(w http.ResponseWriter, r *http.Request) {
	log.Fatal("Not implemented")
	return
}

func getByIdHandler(w http.ResponseWriter, r *http.Request) {
	log.Fatal("Not implemented")
	return
}

func updateByIdHandler(w http.ResponseWriter, r *http.Request) {
	log.Fatal("Not implemented")
	return
}

func deleteByIdHandler(w http.ResponseWriter, r *http.Request) {
	log.Fatal("Not implemented")
	return
}
