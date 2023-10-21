// Package db manages the connection to a PostgreSQL database.
//
// This package provides functionalities to connect to a PostgreSQL database,
// close the database connection, and retrieve the underlying sql.DB object.
package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// Database struct holds the database connection.
type Database struct {
	database *sql.DB
}

func NewDatabase() (*Database, error) {
	db, err := sql.Open("postgres", "postgresql://root:password@localhost:5432/quick-note?sslmode=disable")
	if err != nil {
		return nil, err
	}
	return &Database{database: db}, nil
}

func (d *Database) Close() {
	d.database.Close()
}

func (d *Database) GetDB() *sql.DB {
	return d.database
}
