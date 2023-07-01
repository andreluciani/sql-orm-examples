package main

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Quotes struct {
	ID    uint `gorm:"primaryKey"`
	Quote string
}

var (
	initialQuotes = []Quotes{
		{Quote: "Nothing is impossible"},
		{Quote: "If you`re going through hell, keep going"},
		{Quote: "We need much less than we think we need"},
		{Quote: "If things go wrong, don`t go with them"},
		{Quote: "Whatever you are, be a good one"},
	}
)

func main() {
	dsn := "host=localhost port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// Checking if DB exists
	rs := db.Raw("SELECT * FROM pg_database WHERE datname = 'quotes_db';")
	if rs.Error != nil {
		log.Fatal("Raw query failed:", err)
	}

	// If not, create it
	var rec = make(map[string]interface{})
	if rs.Find(rec); len(rec) == 0 {
		if rs := db.Exec("CREATE DATABASE quotes_db;"); rs.Error != nil {
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
	dsn = "host=localhost dbname=quotes_db port=5432 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Quotes{})

	for _, quote := range initialQuotes {
		db.Create(&quote)
	}

	log.Println("Successfully added seed data!")
}
