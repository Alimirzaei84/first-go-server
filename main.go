package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var movies []Movie

func main() {
    router := mux.NewRouter()

    movies = append(movies, Movie{ID: "1", Isbn: "448743", Title: "Movie One", Director: &Director{Firstname: "John", Lastname: "Doe"}})
    movies = append(movies, Movie{ID: "2", Isbn: "847564", Title: "Movie Two", Director: &Director{Firstname: "Steve", Lastname: "Smith"}})



 // Apply middleware
    router.Use(LoggingMiddleware)
    router.Use(SetHeadersMiddleware)

    router.HandleFunc("/movies", getMovies).Methods("GET")
    router.HandleFunc("/movies/{id}", getMovie).Methods("GET")
    router.HandleFunc("/movies", createMovie).Methods("POST")
    router.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
    router.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

    log.Fatal(http.ListenAndServe(":8000", router))
}