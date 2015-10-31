package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func students(w http.ResponseWriter, r *http.Request) {
	students := Students {
		Student{NetID: "jl1424", Name: "Jon Lee", Major: "Computer Science", Year: 2016, Grade: 4, Rating: "A"},
	}

	if err := json.NewEncoder(w).Encode(students); err != nil {
		panic(err)
	}
}

func showStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	netid := vars["netid"]
	fmt.Fprintln(w, "NetID:", netid)
}
