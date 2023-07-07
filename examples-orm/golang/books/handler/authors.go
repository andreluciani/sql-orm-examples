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
		if r.Method == http.MethodPost {
			c.CreateAuthor(w, r)
			return
		}
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
		if r.Method == http.MethodPatch {
			c.UpdateAuthor(w, r)
			return
		}
		if r.Method == http.MethodDelete {
			c.DeleteAuthor(w, r)
			return
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
	})
}

func (c *Controller) ListAuthors(w http.ResponseWriter, r *http.Request) {
	var authors []model.Author
	err := c.db.Preload("Books").Find(&authors).Error
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
	err := c.db.Preload("Books").First(&author, id).Error
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

func (c *Controller) CreateAuthor(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var payload createAuthorPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	author := &model.Author{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
	}
	if err := c.db.Create(&author).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
	}
	result, err := json.Marshal(author)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(result)
}

type createAuthorPayload struct {
	FirstName string
	LastName  string
}

func (c *Controller) UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/authors/"):]
	var author model.Author
	err := c.db.Preload("Books").First(&author, id).Error
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

	defer r.Body.Close()
	var payload updateAuthorPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	if payload.FirstName != "" {
		author.FirstName = payload.FirstName
	}

	if payload.LastName != "" {
		author.LastName = payload.LastName
	}

	if err := c.db.Save(&author).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
	}
	result, err := json.Marshal(author)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(result)
}

type updateAuthorPayload struct {
	FirstName string
	LastName  string
}

func (c *Controller) DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/authors/"):]
	var author = model.Author{}
	err := c.db.Where("id = ?", id).Delete(&author).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
