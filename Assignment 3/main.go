package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	//"gopkg.in/mgo.v2/bson"
	//"log"
	"gopkg.in/mgo.v2"
)

// fuck modularity

// student.go
type Student struct {
	NetID  string `json: "netid"  bson: "netid"`
	Name   string `json: "name"   bson: "name"`
	Major  string `json: "major"  bson: "major"`
	Year   int    `json: "year"   bson: "year"`
	Grade  int    `json: "grade"  bson: "grade"`
	Rating string `json: "rating" bson: "rating"`
}

// router.go + routes.go
type Route struct {
	Name 		string
	Method 		string
	Path		string
	HandlerFunc	http.HandlerFunc
}
type Routes []Route
/*
var routes = Routes {
	Route {
		"Add student", 
		"POST",
		"/Student",
		AddStudent,
	},
	Route {
		"Get student info",
		"GET",
		"/Student/getstudent",
		GetStudent,
	}, 
	Route {
		"Remove student",
		"DELETE",
		"/Student",
		RemoveStudent,
	}, 
	Route {
		"Update info",
		"PUT",
		"/Student",
		UpdateStudent,
	},
}

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	for _, route := range routes {
		router.Methods(route.Method).Path(route.Path).Name(route.Name).HandlerFunc(route.HandlerFunc)
	}

	return router
}
*/

// handlers.go
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "CS417\n")
}

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

    // add to db
    //t := RepoCreateTodo(student)
    //w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    //w.WriteHeader(http.StatusCreated)
    //if err := json.NewEncoder(w).Encode(t); err != nil {
    //    panic(err)
    //}
} 

func (uc *UserController) AddStudent(w http.ResponseWriter, r *http.Request) {
	
} 

func (uc *UserController) GetStudent(w http.ResponseWriter, r *http.Request) {
	
} 

func (uc *UserController) RemoveStudent(w http.ResponseWriter, r *http.Request) {
	
} 

func (uc *UserController) UpdateStudent(w http.ResponseWriter, r *http.Request) {
	
} 

func (uc *UserController) ListAll(w http.ResponseWriter, r *http.Request) {
	
} 

// db.go
func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}

	return s
}

type UserController struct {
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

// main.go
func main () {
	uc := NewUserController(getSession())
	//router := NewRouter()

	router := mux.NewRouter()
	router.HandleFunc("/", Index)
	router.HandleFunc("/Student", uc.AddStudent).Methods("POST")
	router.HandleFunc("/Student/getstudent", uc.GetStudent).Methods("GET")
	router.HandleFunc("/Student", uc.RemoveStudent).Methods("DELETE")
	router.HandleFunc("/Student", uc.UpdateStudent).Methods("PUT")
	router.HandleFunc("/Student/listall", uc.ListAll).Methods("GET")

	http.ListenAndServe(":1234", router)
}