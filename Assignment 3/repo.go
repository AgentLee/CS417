
package main

import "fmt"

var currentId int

var students Students

// Give us some seed data
func init() {
	RepoAddStudent(Student{Name: "Jon", NetID: "jl1424"})
	RepoAddStudent(Student{Name: "Sej"})
	RepoAddStudent(Student{Name: "Nikki"})
}

func RepoFindStudent(netid string) Student {
	for _, t := range students {
		if t.NetID == netid {
			return t
		}
	}
	// return empty Todo if not found
	return Student{}
}

//this is bad, I don't think it passes race condtions
func RepoAddStudent(t Student) Student {
	currentId += 1
	t.Id = currentId
	students = append(students, t)
	return t
}

func RepoDeleteStudent(id int) error {
	for i, t := range students {
		if t.Id == id {
			students = append(students[:i], students[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Todo with id of %d to delete", id)
}