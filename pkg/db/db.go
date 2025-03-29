package db

import (
	"database/sql"
)

var Database *sql.DB

func InitDb(connStr string) error {
	db, err := sql.Open("postgres", connStr)
	Database = db
	return err
}
