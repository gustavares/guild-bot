package domain

type Respawn struct {
	Id   string
	Name string
}

type Claim struct {
	User
	Respawn
	From string
	To   string
}

func (c *Claim) TimeLeft() string {
	return "Not yet implemented"
}
