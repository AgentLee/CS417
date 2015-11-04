package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route {
		"Index",
		"GET",
		"/",
		Index,
	},
	Route {
		"Student home",
		"GET",
		"/Student",
		StudentIndex,
	},
	Route {
		"List all students",
		"GET",
		"/Student/listall",
		ListAll,
	},
	Route {
		"Add a student",
		"POST",
		"/Student",
		AddStudent,
	},
	Route {
		"Get a student's info",
		"GET",
		"/Student/{netid}",
		GetStudent,
	},
	Route {
		"Delete student",
		"DELETE",
		"/Student",
		RemoveStudent,
	},
	Route {
		"Update information",
		"UPDATE",
		"/Student",
		UpdateStudent,
	},
}