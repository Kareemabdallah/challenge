package main

import (
	"encoding/json" // import encoding/json package
	"fmt"           // import formatted I/O
	"io/ioutil"     // import I/O utilty functions
	"log"           // import logging
	"net/http"      // import HTTP implem.

	"github.com/gorilla/mux" // HTTP router library
)

// function for handling incoming request URL to match given paths
func GetArticles(w http.ResponseWriter, r *http.Request) {

	// Open our jsonFile as a byte array
	byteValue, _ := ioutil.ReadFile("db.json")

	// declaring a data string parsing data as raw
	data := json.RawMessage(string(byteValue))

	// defining map function with keyvalue as a string and a value type
	var articles map[string]*json.RawMessage

	// Unmarshal func parses the JSON-encoded data and stores result in a Go value
	// unmarshalling byteArray which contains jsonFile's content and stores it into 'articles'
	err := json.Unmarshal(data, &articles)
	// if it returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	// print articles content
	fmt.Println(articles)

	// encoding our articles array into a JSON string
	json.NewEncoder(w).Encode(articles)
}

func main() {

	// Defining our HTTP request multiplexer
	router := mux.NewRouter()

	// Registing URL path /articles where acquiring all articles
	router.HandleFunc("/articles", GetArticles).Methods("GET")

	// printing Listening on localhost:9000
	log.Println("Listening on localhost:9000")

	// listening on port 9000
	log.Fatal(http.ListenAndServe(":9000", router))
}
