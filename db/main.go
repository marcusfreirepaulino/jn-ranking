package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitializeDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./sqlite.db")

	if err != nil {
		log.Fatal(err)
	}
}
