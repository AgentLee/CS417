package main

import (
	"log"
	"net/http"
	"flag"
	"fmt"
)

func main() {
	urlPtr		:=	flag.String("url", "", "input url")
    methodPtr 	:= 	flag.String("method", "", "create, remove, etc.")
    yearPtr 	:= 	flag.String("year", "", "remove students with this year")
    dataPtr		:= 	flag.String("data", "", "JSON data for new students")

    flag.Parse()

    fmt.Println("url:", *urlPtr)
    fmt.Println("method:", *methodPtr)
    fmt.Println("year:", *yearPtr)
    fmt.Println("data:", *dataPtr)

	router := NewRouter()

	// need to change to 1234
	// 8080 because idk how to kill the running server
	log.Fatal(http.ListenAndServe(":8080", router))
}