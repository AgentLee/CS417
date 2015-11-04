package main

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