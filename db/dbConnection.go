package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func DbConnection() *sql.DB {
	connStr := "postgresql://postgres:kaskas@localhost/urlshortner?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
