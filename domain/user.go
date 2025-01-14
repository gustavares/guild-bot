package domain

// Nome do player
// Id/Nome do discord
// From (tempo que comeca a vez dele)
// To(tempo que termina a vez dele)
// Id do respawn
// Timestamp

type User struct {
	Characters []Character
	DiscordId  string
}
