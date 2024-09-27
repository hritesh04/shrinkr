package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func connect() *sql.DB {

	connectionStr := os.Getenv("DB_CONNSTR")

	db, err := sql.Open("postgres", connectionStr)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func Init(connectionStr string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connectionStr)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	createUserTable := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(60) NOT NULL,
		email VARCHAR(128) NOT NULL UNIQUE,
		password VARCHAR(60) NOT NULL,
		subscription_type VARCHAR(10) NOT NULL CHECK (subscription_type IN ('free', 'premium')) DEFAULT 'free'
	)
	`

	createUrlTable := `
	CREATE TABLE IF NOT EXISTS urls (
		id SERIAL PRIMARY KEY,
		original VARCHAR(255) NOT NULL,
		shortened VARCHAR(255) NOT NULL,
		user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
		rateRemaining INT DEFAULT 100,
		expiry DATE DEFAULT CURRENT_DATE + INTERVAL '1 year',
		isActive BOOLEAN DEFAULT TRUE
		)
		`

	if _, err := db.Exec(createUserTable); err != nil {
		log.Fatal(err)
		return nil, err
	}
	if _, err := db.Exec(createUrlTable); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return connect(), nil
}
