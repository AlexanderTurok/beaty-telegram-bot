package telegram

type Participant struct {
	Uuid        string `db:"uuid"`
	Name        string `db:"name"`
	Photo       string `db:"photo"`
	Description string `db:"description"`
	Likes       int    `db:"likes"`
}
