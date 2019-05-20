package main

import (
	"fmt"
	"encoding/json"
	"net/http"
)

type Student struct {
	ID string
	Name string
	Grade int
}

var students = []Student{
	Student{"E001", "ethan", 21},
    Student{"W001", "wick", 22},
    Student{"B001", "bourne", 23},
    Student{"B002", "bond", 23},
}

func users(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	if r.Method == "POST" {

		var jsonStudents, err = json.Marshal(students)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(jsonStudents)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func user(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	if r.Method == "POST" {
		var id = r.FormValue("id")

		for _, student := range students {
			if student.ID == id {
				var jsonData, err = json.Marshal(student)
				if err != nil {
					http.Error(w, "", http.StatusInternalServerError)
					return
				}

				w.Write(jsonData)
				return
			}
		}

		http.Error(w, "User not found", http.StatusBadRequest)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func main() {
	http.HandleFunc("/users", users)
	http.HandleFunc("/user", user)

	var port = ":8080"
	fmt.Println("starting web server at http://localhost"+port)
	http.ListenAndServe(port, nil)
}