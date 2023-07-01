package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	connStr := "postgresql://localhost/quotes_db?sslmode=disable"
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to ping database:", err)
	}
	log.Println("Connected to the database")
}

func getRandomQuote() (string, error) {
	rows, err := db.Query("SELECT quote FROM quotes ORDER BY RANDOM() LIMIT 1")
	if err != nil {
		return "", err
	}
	defer rows.Close()

	var quote string
	for rows.Next() {
		err := rows.Scan(&quote)
		if err != nil {
			return "", err
		}
	}

	return quote, nil
}

func quoteHandler(w http.ResponseWriter, r *http.Request) {
	quote, err := getRandomQuote()
	if err != nil {
		log.Println("Failed to retrieve quote:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(quote))
}

func main() {
	http.HandleFunc("/quote", quoteHandler)

	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
