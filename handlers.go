package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"math/rand"
	"net/http"
	"strconv"
)

func GetMovies(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func GetMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, movie := range movies {
		if movie.Id == params["id"] {
			json.NewEncoder(w).Encode(movie)

			return
		}
	}
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var movie Movie
	json.NewDecoder(r.Body).Decode(&movie)
	movie.Id = strconv.Itoa(rand.Intn(1_000_000))
	movies = append(movies, movie)

	json.NewEncoder(w).Encode(movie)
}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for idx, movie := range movies {
		if movie.Id != params["id"] {
			continue
		}

		movies = append(movies[:idx], movies[idx+1:]...)

		var movie Movie
		json.NewDecoder(r.Body).Decode(&movie)
		movie.Id = params["id"]
		movies = append(movies, movie)
		json.NewEncoder(w).Encode(movie)

		break
	}
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for idx, movie := range movies {
		if movie.Id == params["id"] {
			movies = append(movies[:idx], movies[idx+1:]...)

			break
		}
	}

	json.NewEncoder(w).Encode(movies)
}
