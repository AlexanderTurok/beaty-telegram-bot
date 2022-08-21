package telegram

type Voter struct {
	Id           int            `db:"id"`
	Uuid         int            `db:"uuid"`
	Participants *[]Participant `db:"participants"`
}
