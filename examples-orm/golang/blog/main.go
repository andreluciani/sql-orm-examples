package main

import (
	"go-book-server/handler"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/authors", handler.Authors)
	http.HandleFunc("/authors/", handler.AuthorsByID)

	http.HandleFunc("/books", handler.Books)
	http.HandleFunc("/books/", handler.BooksByID)

	log.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
