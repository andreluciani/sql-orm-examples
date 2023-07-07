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
			return
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
	})
}

func (c *Controller) BooksByID() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			c.GetBookByID(w, r)
			return
		}
		if r.Method == http.MethodDelete {
			c.DeleteBook(w, r)
			return
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
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

func (c *Controller) CreateBook(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var payload createBookPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	book := &model.Book{
		Title:             payload.Title,
		Description:       payload.Description,
		YearOfPublication: payload.YearOfPublication,
	}
	if err := c.db.Create(&book).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
	}
	result, err := json.Marshal(book)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(result)
}

type createBookPayload struct {
	Title             string
	Description       string
	YearOfPublication int
	AuthorID          int
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/Books/"):]
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Update Book by ID: " + id))
}

func (c *Controller) DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/books/"):]
	var book = model.Book{}
	err := c.db.Where("id = ?", id).Delete(&book).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
