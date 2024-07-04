package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Student struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Class struct {
	Name     string    `json:"name"`
	Teacher  string    `json:"teacher"`
	Students []Student `json:"students"`
}

var class = Class{
	Name:    "Biology",
	Teacher: "A. Lawrence",
	Students: []Student{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
		{ID: 3, Name: "Charlie"},
	},
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /student/{id}", getStudentInfo)

	http.ListenAndServe(":8080", mux)
}

func getStudentInfo(w http.ResponseWriter, r *http.Request) {
	userRole := r.Header.Get("Role")

	if userRole != "teacher" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for _, student := range class.Students {
		if student.ID == id {
			json.NewEncoder(w).Encode(student)
			return
		}
	}
}
