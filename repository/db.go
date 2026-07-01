package repository

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

func NewDB() (*sql.DB, error) {
	return sql.Open("sqlite", "tasks.db")
}

func InitDB(db *sql.DB) error {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS tasks (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	text TEXT NOT NULL
	)
	`)
	return err
}
