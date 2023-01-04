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

func getMoviesHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(movies)
}

func createMovieHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	var movie Movie
	_ = json.NewDecoder(req.Body).Decode(&movie)

	movie.ID = strconv.FormatInt(rand.Int63n(1000000), 10)
	movies = append(movies, movie)
	json.NewEncoder(res).Encode(movie)
}

func getByIdHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(res).Encode(item)
			return
		}
	}

}

func updateByIdHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(req.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(res).Encode(movie)
			return

		}
	}

	log.Fatal("Not implemented")
	return
}

func deleteByIdHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}

	json.NewEncoder(res).Encode(movies)
}
