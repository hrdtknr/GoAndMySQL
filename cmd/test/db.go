package test

import (
	"database/sql"
	"log"
)

func OpenDB(drive, dsn string) *sql.DB {
	db, err := sql.Open(drive, dsn)
	if err != nil {
		log.Fatal("OpenDB failed:", err)
	}
	return db
}

func CloseDB(db *sql.DB) {
	if err := db.Close(); err != nil {
		log.Fatal("CloseDB failed:", err)
	}
}
