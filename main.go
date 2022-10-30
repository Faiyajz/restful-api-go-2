package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Roll struct {
	ID          string `json:"id"`
	ImageNumber string `json:"imageNumber"`
	Name        string `json:"name"`
	Ingredients string `json:"ingredients"`
}

// As we are not using database to store the data,
// we have created a slice below to store the data
var rolls []Roll

// Index
func getRolls(w http.ResponseWriter, r *http.Request) {

}

// Show
func getRoll(w http.ResponseWriter, r *http.Request) {

}

// Create
func createRoll(w http.ResponseWriter, r *http.Request) {

}

// Update
func updateRoll(w http.ResponseWriter, r *http.Request) {

}

// Delete
func deleteRoll(w http.ResponseWriter, r *http.Request) {

}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/sushi", getRolls).Methods("GET")
	router.HandleFunc("/sushi/{id}", getRoll).Methods("GET")
	router.HandleFunc("/sushi", createRoll).Methods("POST")
	router.HandleFunc("/sushi/{id}", updateRoll).Methods("POST")
	router.HandleFunc("/sushi/{id}", deleteRoll).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":5000", router))

}
