package main

import( 
	"fmt"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	//"strconv"
	"github.com/gorilla/mux"
	"log"
	"time"
)

// basically gave up on modularity

// student 
type Student struct {
	Id        	int       `json:"id"`
	NetID       string    `json:"netid"`
	Name      	string    `json:"name"`
	Major      	string    `json:"major"`
	Year        int       `json:"year"`
	Grade       int       `json:"grade"`
	Rating      string    `json:"rating"`
}
type Students []Student

var currentId int
var students Students

// error
type jsonErr struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}

// handlers
func Index(w http.ResponseWriter, r * http.Request) {
	fmt.Fprint(w, "CS417: Distributed Systems\nREST+Go Project\n")
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

	s := RepoAddStudent(student)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(s); err != nil {
		panic(err)
	}
}

func ListAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(students); err != nil {
		panic(err)
	}
}

func RemoveStudent(w http.ResponseWriter, r *http.Request) {
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

	s := RepoDeleteStudent(student.Year)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(s); err != nil {
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

// repo
func RepoFindStudent(netid string) Student {
	for _, t := range students {
		if t.NetID == netid {
			return t
		}
	}

	return Student{}
}

func RepoAddStudent(t Student) Student {
	currentId += 1
	t.Id = currentId
	students = append(students, t)
	return t
}

func RepoDeleteStudent(year int) error {
	for i, t := range students {
		if t.Year == year {
			students = append(students[:i], students[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Todo with id of %d to delete", year)
}

// logger
func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}

// router
func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}

	return router
}

// routes
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}
type Routes []Route

var routes = Routes {
	Route {
		"Index",
		"GET",
		"/",
		Index,
	},
	Route {
		"Add Student",
		"POST",
		"/Student",
		AddStudent,
	},
	Route {
		"Remove Student",
		"DELETE",
		"/Student",
		RemoveStudent,
	},
	Route {
		"List all",
		"GET",
		"/Student/listall",
		ListAll,
	},
	Route {
		"Get student info",
		"GET",
		"/Student/{netid}",
		GetStudent,
	},
}

func main() {
	router := NewRouter()

	log.Fatal(http.ListenAndServe(":1234", router))
}