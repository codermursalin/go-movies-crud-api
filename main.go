package main

import (
	"log"
	"net/http"

	"github.com/codermursalin/go-movies-crud-api/controllers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/movies", controllers.GetMovies).Methods("GET")
	r.HandleFunc("/movie/{id}", controllers.GetMovie).Methods("GET")
	r.HandleFunc("/movies", controllers.CreateMovie).Methods("POST")
	r.HandleFunc("/movies/delete/{id}", controllers.DeleteMovie).Methods("DELETE")
	r.HandleFunc("/movies/update/{id}", controllers.UpdateMovie).Methods("PATCH")
	log.Fatal(http.ListenAndServe(":8080", r))
}
