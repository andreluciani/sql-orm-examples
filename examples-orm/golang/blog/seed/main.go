package main

import (
	"log"

	"go-book-server/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	initialAuthors = []model.Author{
		{FirstName: "John", LastName: "Doe"},
	}

	initialBooks = []model.Book{
		{Title: "John", Description: "Doe", YearOfPublication: 123, AuthorID: 1},
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
