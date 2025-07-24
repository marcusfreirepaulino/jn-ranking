package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func createParticipantsTable(db *sql.DB) {
	sqlStatement := `
		CREATE TABLE IF NOT EXISTS participants (
			slug TEXT PRIMARY KEY,
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
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			slug TEXT NOT NULL,
			release_date TEXT
		);
	`
	_, err := db.Exec(sqlStatement)
	if err != nil {
		log.Fatalf("Failed to create episodes table: %v", err)
	}
	log.Println("Table 'episodes' created.")
}

func createEpisodesParticipantsTable(db *sql.DB) {
	sqlStatement := `
		CREATE TABLE IF NOT EXISTS episodes_participants (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			episode_id INTEGER NOT NULL,
			participant_key INTEGER NOT NULL,
			FOREIGN KEY(episode_id) REFERENCES episodes(id),
			FOREIGN KEY(participant_key) REFERENCES participants(slug)
		);
	`

	_, err := db.Exec(sqlStatement)
	if err != nil {
		log.Fatalf("Failed to create episodes_participants table: %v", err)
	}
	log.Println("Table 'episodes_participants' created.")
}

func main() {
	db, err := sql.Open("sqlite3", "./sqlite.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	createParticipantsTable(db)
	createEpisodesTable(db)
	createEpisodesParticipantsTable(db)
}
