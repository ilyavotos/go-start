package db

import (
	"database/sql"
	"log"
)

var Database *sql.DB

func InitDb(connStr string) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Db error", err)
	}
	Database = db
}
