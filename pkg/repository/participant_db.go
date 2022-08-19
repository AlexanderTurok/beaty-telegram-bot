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

func (p *ParticipantDB) GetParticipant(uuid int) (*telegram.Participant, error) {
	var par telegram.Participant
	query := fmt.Sprintf("SELECT * FROM participants WHERE uuid=%d;", uuid)
	err := p.db.QueryRow(query).Scan(&par.Id, &par.Uuid, &par.Nickname, &par.Photo, &par.Information, &par.Votes)

	return &par, err
}

func (p *ParticipantDB) AddParticipant(uuid int) error {
	query := fmt.Sprintf("INSERT INTO participants (uuid) VALUES (%d);", uuid)
	_, err := p.db.Exec(query)

	return err
}

func (p *ParticipantDB) UpdateParticipant(column, value string, uuid int) error {
	query := fmt.Sprintf("UPDATE participants SET %s='%s' WHERE uuid=%d", column, value, uuid)
	_, err := p.db.Exec(query)

	return err
}

func (p *ParticipantDB) DeleteParticipant(uuid int) error {
	query := fmt.Sprintf("DELETE FROM participants WHERE uuid=%d;", uuid)
	_, err := p.db.Exec(query)

	return err
}
