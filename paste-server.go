package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Person struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	LastName  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}
type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

// func getSpecificSLR(w http.ResponseWriter, req *http.Request) {
// 	params := mux.Vars(req)
// 	for _, item := range people {
// 		if item.ID == params["id"] {
// 			json.NewEncoder(w).Encode(item)
// 			return
// 		}
// 	}
// 	json.NewEncoder(w).Encode(&Person{})

// }

func getSLR(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(people)

}
func addSLR(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var person Person
	_ = json.NewDecoder(req.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}
func delSLR(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
		}
	}
	json.NewEncoder(w).Encode(people)
}

//
//the main function
//
func main() {
	router := mux.NewRouter()
	people = append(people, Person{ID: "1", Firstname: "Nic", LastName: "Raboy", Address: &Address{City: "Dublin", State: "California"}})
	people = append(people, Person{ID: "2", Firstname: "Maria", LastName: "Raboy"})
	router.HandleFunc("/secretLinkRoute", getSLR).Methods("GET")
	router.HandleFunc("/slrAdd/{id}", addSLR).Methods("POST")
	router.HandleFunc("/slrDelete/{i}", delSLR).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
