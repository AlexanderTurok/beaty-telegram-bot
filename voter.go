package telegram

type Voter struct {
	Id                int            `db:"id"`
	Uuid              int            `db:"uuid"`
	LikedParticipants *[]Participant `db:"liked_participants"`
}
