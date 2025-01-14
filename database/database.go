package database

import (
	"database/sql"
	"fmt"
)

type Db struct {
	db *sql.DB
}

func (d *Db) AddCharacterToUser() error {
	defer d.db.Close()

	// TODO: implement
	return fmt.Errorf("not yet implemented")
}

func New() *Db {
	return &Db{
		db: initSqlLite(),
	}
}

func NewWithDSN(dsn string) *Db {
	return &Db{
		db: initSqlLiteWithDSN(dsn),
	}
}
