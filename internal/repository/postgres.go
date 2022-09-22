package repository

import (
	"github.com/jmoiron/sqlx"
)

const (
	participantTable       = "participant"
	voterTable             = "voter"
	votersParticipantTable = "voter_participant"
)

func NewPostgresDB(config Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", "dbname=tgbeaty_db sslmode=disable")
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
