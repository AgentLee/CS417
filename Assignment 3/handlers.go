package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "CS417: Distributed Systems\n")
	fmt.Fprint(w, "Assignment 3: REST and Go")
}

func StudentIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "/Student/listall\n")
	fmt.Fprint(w, "/Student/{netid}")
}

func ListAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(students); err != nil {
		panic(err)
	}
}

func GetStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var todoId string = vars["netid"]
	student := RepoFindStudent(todoId)
	if student.NetID != "" {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(student); err != nil {
			panic(err)
		}
		return
	}

	// If we didn't find it, 404
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}

}

/*
Test with this curl command:

curl -H "Content-Type: application/json" -d '{"name":"New Todo"}' http://localhost:8080/todos

*/
func AddStudent(w http.ResponseWriter, r *http.Request) {
	var student Student
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &student); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	s := RepoAddStudent(student)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(s); err != nil {
		panic(err)
	}
}

func UpdateStudent (w http.ResponseWriter, r *http.Request) {

}

func RemoveStudent (w http.ResponseWriter, r *http.Request) {

}