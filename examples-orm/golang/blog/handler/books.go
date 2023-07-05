package handler

import (
	"encoding/json"
	"errors"
	"go-book-server/model"
	"log"
	"net/http"

	"gorm.io/gorm"
)

func (c *Controller) Books() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			c.ListBooks(w, r)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Method not allowed"))
		}
	})
}

func (c *Controller) BooksByID() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			c.GetBookByID(w, r)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Method not allowed"))
		}
	})
}

func (c *Controller) ListBooks(w http.ResponseWriter, r *http.Request) {
	var Books []model.Book
	err := c.db.Preload("Author").Find(&Books).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Book not found."))
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
		return
	}
	result, err := json.Marshal(Books)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func (c *Controller) GetBookByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/Books/"):]
	var Book model.Book
	err := c.db.Preload("Author").First(&Book, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Book not found."))
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
		return
	}
	result, err := json.Marshal(Book)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Create Book"))
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/Books/"):]
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Update Book by ID: " + id))
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/Books/"):]
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Delete Book by ID: " + id))
}
