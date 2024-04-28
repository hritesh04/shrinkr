package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func Connect() *sql.DB{

	connectionStr := os.Getenv("DB_CONNSTR")

	db, err := sql.Open("postgres", connectionStr)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func Init() error{

	connectionStr := os.Getenv("DB_CONNSTR")

	db,err := sql.Open("postgres",connectionStr)
	if err!= nil{
		log.Fatal(err)
		return err
	}

	createUserTable := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(60) NOT NULL,
		email VARCHAR(128) NOT NULL UNIQUE,
		password VARCHAR(60) NOT NULL
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
		rateLimitReset DATE DEFAULT CURRENT_DATE + INTERVAL '1 day',
		isActive BOOLEAN DEFAULT TRUE
		)
		`
		
		if _,err:=db.Exec(createUserTable);err != nil {
			log.Fatal(err)
			return err
		}
		if _,err:=db.Exec(createUrlTable);err != nil {
			log.Fatal(err)
			return err
		}
	return nil
}