package repository

import (
	"database/sql"
	"fmt"

	"github.com/AlexanderTurok/telegram-beaty-bot"
)

type VoterDB struct {
	db *sql.DB
}

func NewVoterDB(db *sql.DB) *VoterDB {
	return &VoterDB{
		db: db,
	}
}

func (v *VoterDB) GetParticipant(uuid string) (*telegram.Participant, error) {
	var p telegram.Participant
	query := fmt.Sprintf("SELECT * FROM participants WHERE uuid=%s;", uuid)
	err := v.db.QueryRow(query).Scan(&p.Id, &p.Uuid, &p.Nickname, &p.Photo, &p.Information, &p.Votes)

	return &p, err
}

func (v *VoterDB) UpdateParticipant(column, value, uuid string) error {
	query := fmt.Sprintf("UPDATE participants SET %s='%s' WHERE uuid=%s", column, value, uuid)
	_, err := v.db.Exec(query)

	return err
}