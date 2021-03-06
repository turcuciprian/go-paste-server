package main

import (
	"encoding/json"
	"log"
	"net/http"

	"fmt"

	"github.com/gorilla/mux"
)

type codeData struct {
	ID         string `json:"id,omitempty"`
	SecretCode string `json:"sc,omitempty"`
}

var codes []codeData

// func getSpecificSLR(w http.ResponseWriter, req *http.Request) {
// 	params := mux.Vars(req)
// 	for _, item := range codes {
// 		if item.ID == params["id"] {
// 			json.NewEncoder(w).Encode(item)
// 			return
// 		}
// 	}
// 	json.NewEncoder(w).Encode(&codeData{})

// }

func getSLR(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(codes)

}
func modSLR(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var jsonResponse codeData
	_ = json.NewDecoder(req.Body).Decode(&jsonResponse)
	jsonResponse.ID = params["id"]
	codes[0] = jsonResponse
	json.NewEncoder(w).Encode(codes)
}

func delSLR(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range codes {
		fmt.Println("item.ID = ", item.ID)
		fmt.Println("params[id] = ", params["id"])

		if item.ID == params["id"] {
			codes = append(codes[:index], codes[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(codes)
}

//
//the main function
//
func main() {
	router := mux.NewRouter()
	//one and only code
	codes = append(codes, codeData{ID: "1", SecretCode: "Empty"})

	router.HandleFunc("/slr", getSLR).Methods("GET")
	router.HandleFunc("/slr", modSLR).Methods("POST")
	router.HandleFunc("/slr/{id}", delSLR).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
