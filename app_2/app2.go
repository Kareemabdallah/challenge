package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	id      string `json:"id"`
	message string `json:"message"`
}

func handler(w http.ResponseWriter, r *http.Request) {

	resp, err := http.Get("http://localhost:9000/")

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close() //defering for future usage

	var result map[string]interface{} //parsing JSON strings

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal([]byte(body), &result) //Uunmarshalling content into result

	str := fmt.Sprint(result["message"]) //converting data type interface to string
	str1 := Reverse(str)
	fmt.Fprintf(w, str1)

}

func Reverse(s string) string { // converts strings to rune slices
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 { //
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", handler).Methods("GET")
	log.Println("Listening on localhost:7000")
	log.Fatal(http.ListenAndServe(":7000", r))

}
