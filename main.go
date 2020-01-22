package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Articles struct {
	Articles []Article `json:"articles"`
}

type Article struct {
	id      string `json:"id"`
	message string `json:"message"`
}

func GetArticles(w http.ResponseWriter, r *http.Request) {

	byteValue, _ := ioutil.ReadFile("db.json")
	data := json.RawMessage(string(byteValue)) // declaring a data string parsing data as raw
	json.NewEncoder(w).Encode(&data)

	var generic map[string]interface{}
	err := json.Unmarshal([]byte(byteValue), &generic)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(generic["articles"])

}

func main() {

	router := mux.NewRouter() // HTTP request multiplexer
	router.HandleFunc("/articles", GetArticles).Methods("GET")
	log.Println("Listening on localhost:9000")
	log.Fatal(http.ListenAndServe(":9000", router))
}
