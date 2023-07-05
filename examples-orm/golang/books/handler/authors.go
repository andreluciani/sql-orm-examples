package handler

import (
	"encoding/json"
	"errors"
	"go-book-server/model"
	"log"
	"net/http"

	"gorm.io/gorm"
)

func (c *Controller) Authors() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			c.ListAuthors(w, r)
			return
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
	})
}

func (c *Controller) AuthorsByID() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			c.GetAuthorByID(w, r)
			return
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
	})
}

func (c *Controller) ListAuthors(w http.ResponseWriter, r *http.Request) {
	var authors []model.Author
	err := c.db.Find(&authors).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("author not found."))
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
		return
	}
	result, err := json.Marshal(authors)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func (c *Controller) GetAuthorByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/authors/"):]
	var author model.Author
	err := c.db.First(&author, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("author not found."))
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
		return
	}
	result, err := json.Marshal(author)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(result)
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
