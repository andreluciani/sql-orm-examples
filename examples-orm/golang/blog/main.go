package main

import (
	"go-book-server/handler"
	"log"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	dsn := "host=localhost dbname=books_db port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	controller := handler.NewController(db)

	http.HandleFunc("/authors", controller.Authors())
	http.HandleFunc("/authors/", controller.AuthorsByID())

	log.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
