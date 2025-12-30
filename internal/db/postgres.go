package database

import (
	"database/sql"
	"os"
	"time"
)

func ConnectToDB() (*sql.DB, error) {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == " " {
		panic("DATABASE_URL is not set")
	}
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	if err := db.Ping(); err != nil {
		panic(err)
	}
	return db, nil
}
