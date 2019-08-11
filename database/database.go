package database

import (
	"database/sql"
	"log"
)

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(172.17.0.2:3306)/todo")

	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Database connected")
	}

	return db
}
