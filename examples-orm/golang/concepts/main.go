package main

import (
	"log"
	"math/rand"
	"net/http"
)

var quotes = []string{
	"Nothing is impossible.\n",
	"If you're going through hell, keep going.\n",
	"We need much less than we think we need.\n",
	"If things go wrong, don't go with them.\n",
	"Whatever you are, be a good one.\n",
}

func quoteHandler(w http.ResponseWriter, r *http.Request) {
	index := rand.Intn(len(quotes))
	quote := quotes[index]
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(quote))
}

func main() {
	http.HandleFunc("/quote", quoteHandler)

	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
