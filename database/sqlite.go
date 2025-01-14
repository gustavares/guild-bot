package database

import (
	"database/sql"
	"log"

	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
)

const DEFAULT_DSN = "file:/root/db/guild-bot.db"

func initSqlLite() *sql.DB {
	return initSqlLiteWithDSN(DEFAULT_DSN)
}

func initSqlLiteWithDSN(dsn string) *sql.DB {
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	log.Println("Connected to Database!")
	return db
}
