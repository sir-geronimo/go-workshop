package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

var (
	movies []Movie
)

func main() {
	initMovies()

	r := mux.NewRouter()

	r.HandleFunc("/movies", GetMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", GetMovie).Methods("GET")
	r.HandleFunc("/movies", CreateMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", UpdateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", DeleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port: 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}

func initMovies() {
	movies = append(movies, Movie{
		Id:    strconv.Itoa(rand.Intn(1_000_000)),
		Isbn:  "123",
		Title: "Movie one",
		Director: &Director{
			Id:        strconv.Itoa(rand.Intn(1_000_000)),
			Firstname: "John",
			Lastname:  "McCarter",
		},
	})

	movies = append(movies, Movie{
		Id:    strconv.Itoa(rand.Intn(1_000_000)),
		Isbn:  "321",
		Title: "Movie two",
		Director: &Director{
			Id:        strconv.Itoa(rand.Intn(1_000_000)),
			Firstname: "Marie",
			Lastname:  "Jane",
		},
	})
}
