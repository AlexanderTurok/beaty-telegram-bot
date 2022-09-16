package repository

import (
	"context"
	"fmt"

	telegram "github.com/AlexanderTurok/telegram-beaty-bot/pkg"
	"github.com/go-redis/redis/v9"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type VoterRepository struct {
	context context.Context
	db      *sqlx.DB
	redis   *redis.Client
}

func NewVoterRepository(context context.Context, db *sqlx.DB, redis *redis.Client) *VoterRepository {
	return &VoterRepository{
		context: context,
		db:      db,
		redis:   redis,
	}
}

func (r *VoterRepository) Create(uuid int64) error {
	query := fmt.Sprintf("INSERT INTO %s (uuid) VALUES ($1)", voterTable)
	_, err := r.db.Exec(query, uuid)
	if err != nil {
		return err
	}

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

func (r *VoterRepository) GetParticipant(uuid int64) (telegram.Participant, error) {
	var participant telegram.Participant

	query := fmt.Sprintf("SELECT participant_uuid FROM %s LEFT JOIN %s ON participant.uuid = voter_participant.participant_uuid WHERE voter_participant.voter_uuid = $1 LIMIT 1",
		votersParticipantTable, participantTable)
	err := r.db.Get(&participant, query, uuid)

	return participant, err
}

func (r *VoterRepository) DeleteParticipant(voteUuid int64, participantUuid string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE voter_uuid = $1 AND participant_uuid = $2",
		votersParticipantTable)
	_, err := r.db.Exec(query)

	return err
}
