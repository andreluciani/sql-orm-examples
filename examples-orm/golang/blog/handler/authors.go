package handler

import (
	"net/http"
)

func Authors(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		ListAuthors(w, r)
	} else if r.Method == http.MethodPost {
		Create(w, r)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
	}
}

func AuthorsByID(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		GetAuthorByID(w, r)
	} else if r.Method == http.MethodPatch {
		UpdateAuthor(w, r)
	} else if r.Method == http.MethodDelete {
		DeleteAuthor(w, r)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
	}
}

func ListAuthors(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Get all authors"))
}

func GetAuthorByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/authors/"):]
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Get author by ID: " + id))
}

func CreateAuthor(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Create author"))
}

func UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/authors/"):]
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Update author by ID: " + id))
}

func DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/authors/"):]
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Delete author by ID: " + id))
}
