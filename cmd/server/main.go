package main

import (
	"Assignment3/internal/app/db"
	"fmt"
	"log"
)

func main() {
	database, err := db.NewSQLite("test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	_, err = database.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			email TEXT NOT NULL
		);
	`)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("SQLite database is working")
}
