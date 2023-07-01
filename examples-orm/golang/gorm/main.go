package main

import (
	"log"
	"net/http"
)

func quoteHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Nothing is impossible.\n"))
}

func main() {
	http.HandleFunc("/quote", quoteHandler)

	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
