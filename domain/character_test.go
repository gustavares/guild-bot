package domain

import (
	"guild-bot/database"
	"testing"
)

func TestAddCharacterToUser(t *testing.T) {
	db := database.NewWithDSN("file:memdb?mode=memory&cache=shared")
	err := AddCharacterToUser(db, "TestName", "123456")

	if err != nil {
		t.Errorf("got %s want nil", err.Error())
	}
}
