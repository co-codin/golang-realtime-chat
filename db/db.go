package db

import "database/sql"

type Database struct {
	db *sql.DB
}

func NewDatabase(db *sql.DB) (*Database, error) {
	db, err := sql.Open("postgres", "")
}
