package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func createParticipantsTable(db *sql.DB) {
	sqlStatement := `
		CREATE TABLE IF NOT EXISTS participants (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT
			name TEXT NOT NULL
		);
	`

	_, err := db.Exec(sqlStatement)
	if err != nil {
		log.Fatalf("Failed to create episodes table: %v", err)
	}
	log.Println("Table 'participants' created.")
}

func createEpisodesTable(db *sql.DB) {
	sqlStatement := `
		CREATE TABLE IF NOT EXISTS episodes (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL
			slug TEXT NOT NULL
			release_date TEXT
		);
	`
	_, err := db.Exec(sqlStatement)
	if err != nil {
		log.Fatalf("Failed to create episodes table: %v", err)
	}
	log.Println("Table 'episodes' created.")
}

func main() {
	db, err := sql.Open("sqlite3", "../db/sqlite.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	createParticipantsTable(db)

	createEpisodesTable(db)
}
