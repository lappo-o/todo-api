package repository

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func NewDB() (*sql.DB, error) {
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		return nil, fmt.Errorf("DB_URL is not set")
	}
	return sql.Open("pgx", dbURL)
}

func InitDB(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			email TEXT NOT NULL UNIQUE,	
			password TEXT NOT NULL
		)
	`)
	if err != nil {
		return err
	}
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS tasks (
			user_id INTEGER NOT NULL REFERENCES users(id),
			id SERIAL PRIMARY KEY,
			text TEXT NOT NULL
		)
	`)

	return err
}
