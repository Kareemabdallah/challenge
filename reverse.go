package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Articles struct {
	Articles []Article `json:"articles"`
}

type Article struct {
	id      string `json:"id"`
	message string `json:"message"`
}

func main() {
	resp, err := http.Get("http://localhost:9000/articles")
	if err != nil {
		log.Fatal(err)
	}
	var generic map[string]interface{}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal([]byte(body), &generic)

	str := fmt.Sprint(generic["articles"]) //converting data type interface to string
	str1 := Reverse(str)
	fmt.Println(str1)
}

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
