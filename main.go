package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

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
	w.Header().Set("Content-Type", "application/json") //Set the headers and the response
	json.NewEncoder(w).Encode(rolls)                   //render our rolls slice as json and send it to the response stream.
}

// Show
func getRoll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r) //mux.Vars() function is setting our params variable from the http response we are passing it.
	for _, item := range rolls {

		//when we find the item in our slice where the ID matches the id being sent
		//through our params variable we should render it as json just like we did
		//in the getRolls() function and then return.
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

// Create
func createRoll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newRoll Roll //create new instance of the Struct Roll

	json.NewDecoder(r.Body).Decode(&newRoll)
	//read data from our requests by passing the body of our http request e.g. json.NewDecoder(r.Body)
	//Call .Decode() passing it a pointer to our newRoll Struct which is an instance of Roll Struct
	//which allows it to match the json to the appropriate properties of the struct

	newRoll.ID = strconv.Itoa(len(rolls) + 1) //convert int to string

	rolls = append(rolls, newRoll) //add the newRoll to the rolls slice

	json.NewEncoder(w).Encode(newRoll) //sending a response back containing our newRoll
}

// Update
func updateRoll(w http.ResponseWriter, r *http.Request) {

}

// Delete
func deleteRoll(w http.ResponseWriter, r *http.Request) {

}

func main() {

	rolls = append(rolls, Roll{ID: "1", ImageNumber: "8", Name: "Spicy Tuna Roll", Ingredients: "Tuna, Chili sauce, Nori, Rice"})

	router := mux.NewRouter()

	router.HandleFunc("/sushi", getRolls).Methods("GET")
	router.HandleFunc("/sushi/{id}", getRoll).Methods("GET")
	router.HandleFunc("/sushi", createRoll).Methods("POST")
	router.HandleFunc("/sushi/{id}", updateRoll).Methods("POST")
	router.HandleFunc("/sushi/{id}", deleteRoll).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":5000", router))

}
