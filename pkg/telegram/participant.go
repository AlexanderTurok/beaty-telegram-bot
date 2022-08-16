package telegram

type Participant struct {
	Id          int         `db:"id"`
	Uuid        int         `db:"uuid"`
	Nickname    string      `db:"nickname"`
	Photo       string      `db:"photo"`
	Information string      `db:"information"`
	Votes       interface{} `db:"votes"`
}
