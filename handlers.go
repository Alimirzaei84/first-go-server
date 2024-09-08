package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func getMovies(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for _, item := range movies {
        if item.ID == params["id"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    json.NewEncoder(w).Encode(&Movie{})
}

func createMovie(w http.ResponseWriter, r *http.Request) {
    var movie Movie
    if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    movies = append(movies, movie)
    json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for index, item := range movies {
        if item.ID == params["id"] {
            movies = append(movies[:index], movies[index+1:]...)
            var movie Movie
            if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
                http.Error(w, err.Error(), http.StatusBadRequest)
                return
            }
            movie.ID = params["id"]
            movies = append(movies, movie)
            json.NewEncoder(w).Encode(movie)
            return
        }
    }
    json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for index, item := range movies {
        if item.ID == params["id"] {
            movies = append(movies[:index], movies[index+1:]...)
            break
        }
    }
    json.NewEncoder(w).Encode(movies)
}