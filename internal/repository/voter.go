package repository

import (
	"context"
	"database/sql"
	"fmt"

	telegram "github.com/AlexanderTurok/telegram-beaty-bot/pkg"
	"github.com/go-redis/redis/v9"
	_ "github.com/lib/pq"
)

type VoterRepository struct {
	context context.Context
	db      *sql.DB
	redis   *redis.Client
}

func NewVoterRepository(context context.Context, db *sql.DB, redis *redis.Client) *VoterRepository {
	return &VoterRepository{
		context: context,
		db:      db,
		redis:   redis,
	}
}

// FIXME:
func (r *VoterRepository) Create(uuid int64) error {
	query := fmt.Sprintf("INSERT INTO %s (uuid) VALUES ($1)", voterTable)
	_, err := r.db.Exec(query, uuid)
	if err != nil {
		return err
	}

	query = fmt.Sprintf("INSERT INTO %s (voter_uuid, participant_uuid) VALUES ($1, (SELECT uuid FROM %s))",
		votersParticipantTable, participantTable)
	_, err = r.db.Exec(query)

	return err
}

func (r *VoterRepository) IsExists(uuid int64) (bool, error) {
	var exists bool
	query := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM %s WHERE uuid=$1)", voterTable)
	err := r.db.QueryRow(query, uuid).Scan(&exists)

	return exists, err
}

func (r *VoterRepository) Activate(uuid int64) error {
	query := fmt.Sprintf("INSERT INTO %s (voter_uuid, participant_uuid) SELECT $1, uuid FROM %s WHERE participant.uuid <> $1",
		votersParticipantTable, participantTable)
	_, err := r.db.Exec(query, uuid)

	return err
}

// FIXME:
func (r *VoterRepository) GetParticipant(uuid int64) (telegram.Participant, error) {
	// get participants uuid
	// get participant by id
	// delete participant uuid from voters_participant
	return telegram.Participant{}, nil
}
