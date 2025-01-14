package domain

import (
	"guild-bot/database"
)

type Character struct {
	Name string
}

func AddCharacterToUser(db *database.Db, name string, discordId string) error {
	// TODO: implement
	return db.AddCharacterToUser()
}
