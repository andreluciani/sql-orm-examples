package main

import (
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	response := "Hello, World!"

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(response))
}

func main() {
	http.HandleFunc("/hello", helloHandler)

	log.Println("Server started on http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
