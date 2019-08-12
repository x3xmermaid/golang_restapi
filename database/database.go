package database

import (
	"database/sql"
	"log"
)

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/todo")

	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Database connected")
	}

	return db
}
