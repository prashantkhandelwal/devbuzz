package db

import (
	"database/sql"
	"errors"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitDB() {

	if _, err := os.Stat("feeds.db"); errors.Is(err, os.ErrNotExist) {
		d, err := os.Create("feeds.db")
		if err != nil {
			log.Fatalf("ERROR: Cannot create database: %v", err)
		}

		log.Printf("DB: Created - feeds.db")

		d.Close()
	}

	sqldb, err := sql.Open("sqlite3", "feeds.db")
	if err != nil {
		log.Fatalf("ERROR: Cannot open database %v", err)
	}

	db = sqldb

	createFeedTable(sqldb)

	defer sqldb.Close()

	log.Println("DB Setup Complete")
}

// Create a table to store feed URLs
func createFeedTable(db *sql.DB) {
	log.Println("Creating table: feeds")
	query := `CREATE TABLE IF NOT EXISTS feeds (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		'Title'	TEXT,
		'Link'  TEXT,
		'Description'	TEXT,
		'Type'	TEXT,
		'Version' REAL);`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("ERROR: Cannot create feeds table %v", err)
	}

	db.Close()
}

// func Exec(query string) error {

// }

// func AddFeed(db *sql.DB) (bool, error) {

// }
