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

func GetPersonEndPoint(w http.ResponseWriter, req *http.Request) {

}

func GetPeopleEndPoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(people)

}
func CreatePersonEndPoint(w http.ResponseWriter, req *http.Request) {

}
func DeletePersonEndPoint(w http.ResponseWriter, req *http.Request) {

}

//
//the main function
//
func main() {
	router := mux.NewRouter()
	people = append(people, Person{ID: "1", Firstname: "Nic", LastName: "Raboy", Address: &Address{City: "Dublin", State: "California"}})
	people = append(people, Person{ID: "2", Firstname: "Maria", LastName: "Raboy"})
	router.HandleFunc("/people", GetPeopleEndPoint).Methods("GET")
	router.HandleFunc("/people/{id}", GetPersonEndPoint).Methods("GET")
	router.HandleFunc("/people/{id", CreatePersonEndPoint).Methods("POST")
	router.HandleFunc("/people/{id", DeletePersonEndPoint).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
