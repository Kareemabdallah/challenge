package main

import (
	"encoding/json" // import encoding/json package
	"fmt"           // import formatted I/O
	"log"           // import logging
	"net/http"      // import HTTP implem.
)

func main() {
	resp, err := http.Get("http://localhost:9000/articles")
	if err != nil {
		log.Fatal(err)
	}
	var generic map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&generic)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(generic)
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
