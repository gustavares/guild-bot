package database

import (
	"database/sql"
	"log"

	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
)

func Init() *sql.DB {
	db, err := sql.Open("sqlite3", "file:/root/db/guild-bot.db")
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	log.Println("Connected to Database!")
	return db
}

func GetAllRespawns(db *sql.DB) {
	// TODO: implement
}
