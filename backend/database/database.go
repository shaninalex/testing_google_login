package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Database struct {
	DB *sql.DB
}

func CreateConnection(connection_string string) (*Database, error) {
	log.Println("Attempt to create connection")
	var database Database
	db, err := sql.Open("postgres", connection_string)
	if err != nil {
		return nil, err
	}
	log.Println("Database connection established")
	database.DB = db
	return &database, nil
}
