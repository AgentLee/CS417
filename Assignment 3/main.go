package main

import (
	"log"
	"net/http"
)

func main() {
	router := NewRouter()

	// need to change to 1234
	// 8080 because idk how to kill the running server
	log.Fatal(http.ListenAndServe(":8080", router))
}