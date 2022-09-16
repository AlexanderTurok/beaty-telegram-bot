package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	participantTable       = "participant"
	voterTable             = "voter"
	votersParticipantTable = "voter_participant"
)

func NewPostgresDB(config Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		config.Host, config.Port, config.Username, config.DBName, config.Password, config.SSLMode))
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
