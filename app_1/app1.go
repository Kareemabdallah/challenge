package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Reading JSON string and encoding data
func getarticles(w http.ResponseWriter, r *http.Request) {

	byteValue, _ := ioutil.ReadFile("static/db.json")
	data := json.RawMessage(string(byteValue)) // declaring a data string parsing data as raw
	json.NewEncoder(w).Encode(&data)           //Encoding JSON to Server response
	fmt.Println(data)

}

func main() {

	router := mux.NewRouter() // HTTP request multiplexer
	router.HandleFunc("/", getarticles).Methods("GET")
	log.Println("Listening on localhost:9000")
	log.Fatal(http.ListenAndServe(":9000", router))
}
