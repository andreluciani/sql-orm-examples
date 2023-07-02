package handler

import (
	"net/http"
)

func Books(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		List(w, r)
	} else if r.Method == http.MethodPost {
		Create(w, r)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
	}
}

func BooksByID(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		GetByID(w, r)
	} else if r.Method == http.MethodPatch {
		Update(w, r)
	} else if r.Method == http.MethodDelete {
		Delete(w, r)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
	}
}

func List(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Get all books"))
}

func GetByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/books/"):]
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Get book by ID: " + id))
}

func Create(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Create book"))
}

func Update(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/books/"):]
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Update book by ID: " + id))
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/books/"):]
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Delete book by ID: " + id))
}
