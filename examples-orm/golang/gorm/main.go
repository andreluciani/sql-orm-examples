package main

import (
	"log"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type Quotes struct {
	ID    uint `gorm:"primaryKey"`
	Quote string
}

func getRandomQuote() (string, error) {
	var quote Quotes
	err := db.Order("RANDOM()").Take(&quote).Error
	if err != nil {
		return "", err
	}
	return quote.Quote, nil
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
	dsn := "host=localhost dbname=quotes_db port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	http.HandleFunc("/quote", quoteHandler)

	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
