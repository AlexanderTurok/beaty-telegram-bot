package repository

import (
	"database/sql"
	"fmt"

	"github.com/AlexanderTurok/telegram-beaty-bot"
)

type ParticipantDB struct {
	db *sql.DB
}

func NewParticipantDB(db *sql.DB) *ParticipantDB {
	return &ParticipantDB{
		db: db,
	}
}

func (p *ParticipantDB) GetParticipant(uuid string) (*telegram.Participant, error) {
	var par telegram.Participant
	query := fmt.Sprintf("SELECT * FROM participants WHERE uuid=%s;", uuid)
	err := p.db.QueryRow(query).Scan(&par.Id, &par.Uuid, &par.Nickname, &par.Photo, &par.Information, &par.Votes)

	return &par, err
}

func (p *ParticipantDB) AddParticipant(uuid string) error {
	query := fmt.Sprintf("INSERT INTO participants (uuid) VALUES (%s);", uuid)
	_, err := p.db.Exec(query)

	return err
}

func (p *ParticipantDB) UpdateParticipant(column, value, uuid string) error {
	query := fmt.Sprintf("UPDATE participants SET %s='%s' WHERE uuid=%s", column, value, uuid)
	_, err := p.db.Exec(query)

	return err
}

func (p *ParticipantDB) DeleteParticipant(uuid string) error {
	query := fmt.Sprintf("DELETE FROM participants WHERE uuid=%s;", uuid)
	_, err := p.db.Exec(query)

	return err
}
