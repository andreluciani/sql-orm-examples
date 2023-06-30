package main

import (
	"database/sql"
	"log"
	"math/rand"
	"net/http"

	_ "github.com/lib/pq"
)

var quotes = []string{
	"Nothing is impossible.\n",
	"If you're going through hell, keep going.\n",
	"We need much less than we think we need.\n",
	"If things go wrong, don't go with them.\n",
	"Whatever you are, be a good one.\n",
}

var db *sql.DB

func init() {
	// Connect to the PostgreSQL database
	connStr := "postgres://localhost:5432/quotes_db?sslmode=disable"
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Set the maximum number of idle connections in the connection pool
	db.SetMaxIdleConns(5)

	// Set the maximum number of open connections to the database
	db.SetMaxOpenConns(10)

	// Ping the database to verify the connection
	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	log.Println("Connected to the database")
}

func getRandomQuote() (string, error) {
	// Query a random quote from the "quotes" table
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
	quote_db, err := getRandomQuote()
	if err != nil {
		log.Println("Failed to retrieve quote:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	log.Println("Retrieved quote (DB):", quote_db)

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
