package main

import (
	"log"

	"go-book-server/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	initialAuthors = []model.Author{
		{FirstName: "William", LastName: "Shakespeare"},
		{FirstName: "Harper", LastName: "Lee"},
	}

	initialBooks = []model.Book{
		{
			Title:             "Macbeth",
			Description:       "A Scottish general's ruthless quest for power leads to his descent into madness and bloodshed, exposing the tragic consequences of unchecked ambition.",
			YearOfPublication: 1600,
			AuthorID:          1,
		},
		{
			Title:             "Romeo and Juliet",
			Description:       " The forbidden love between two young individuals from warring families in Verona ends in tragedy, exploring themes of love, fate, and the destructive power of hatred.",
			YearOfPublication: 1595,
			AuthorID:          1,
		},
		{
			Title:             "To Kill a Mockingbird",
			Description:       "Set in the racially-charged 1930s Deep South, a young girl's perspective reveals the profound impact of racial injustice and the pursuit of justice through her father's defense of an innocent black man.",
			YearOfPublication: 1860,
			AuthorID:          2,
		},
	}
)

func main() {
	dsn := "host=localhost port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// Checking if DB exists
	rs := db.Raw("SELECT * FROM pg_database WHERE datname = 'books_db';")
	if rs.Error != nil {
		log.Fatal("Raw query failed:", err)
	}

	// If not, create it
	var rec = make(map[string]interface{})
	if rs.Find(rec); len(rec) == 0 {
		if rs := db.Exec("CREATE DATABASE books_db;"); rs.Error != nil {
			log.Fatal("Couldn't create database: ", err)
		}

		// Close db connection
		sql, err := db.DB()
		defer func() {
			_ = sql.Close()
		}()
		if err != nil {
			log.Fatal("An error occurred: ", err)
		}
	}

	// Reconnect and add initial data
	dsn = "host=localhost dbname=books_db port=5432 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&model.Author{}, &model.Book{})

	for _, author := range initialAuthors {
		db.Create(&author)
	}
	for _, book := range initialBooks {
		db.Create(&book)
	}

	log.Println("Successfully added seed data!")
}
