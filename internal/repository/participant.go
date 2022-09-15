package repository

import (
	"context"
	"database/sql"
	"fmt"

	telegram "github.com/AlexanderTurok/telegram-beaty-bot/pkg"
	"github.com/go-redis/redis/v9"
	_ "github.com/lib/pq"
)

type ParticipantRepository struct {
	context context.Context
	db      *sql.DB
	redis   *redis.Client
}

func NewParticipantRepository(context context.Context, db *sql.DB, redis *redis.Client) *ParticipantRepository {
	return &ParticipantRepository{
		context: context,
		db:      db,
		redis:   redis,
	}
}

func (r *ParticipantRepository) Register(uuid int64) error {
	query := fmt.Sprintf("INSERT INTO %s (uuid) VALUES ($1)", participantTable)
	_, err := r.db.Exec(query, uuid)

	return err
}

func (r *ParticipantRepository) IsExists(uuid int64) (bool, error) {
	var exists bool
	query := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM %s WHERE uuid=$1)", participantTable)
	err := r.db.QueryRow(query, uuid).Scan(&exists)

	return exists, err
}

func (r *ParticipantRepository) Activate(uuid int64) error {
	query := fmt.Sprintf("INSERT INTO %s (participant_uuid, voter_uuid) SELECT $1, uuid FROM %s WHERE voter.uuid <> $1",
		votersParticipantTable, voterTable)
	_, err := r.db.Exec(query, uuid)

	return err
}

func (r *ParticipantRepository) Get(uuid int64) (telegram.Participant, error) {
	var participant telegram.Participant
	query := fmt.Sprintf("SELECT * FROM %s WHERE uuid=$1", participantTable)
	err := r.db.QueryRow(query, uuid).Scan(
		&participant.Uuid,
		&participant.Name,
		&participant.Photo,
		&participant.Description,
		&participant.Likes,
	)

	return participant, err
}

func (r *ParticipantRepository) Delete(uuid int64) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE uuid=$1;", participantTable)
	_, err := r.db.Exec(query, uuid)

	return err
}

func (r *ParticipantRepository) GetName(uuid, name string) error {
	err := r.redis.Set(r.context, uuid, name, 0)
	return err.Err()
}

func (r *ParticipantRepository) GetPhoto(uuid, photo string) error {
	err := r.redis.Set(r.context, uuid, photo, 0)
	return err.Err()
}

func (r *ParticipantRepository) GetDescription(uuid, description string) error {
	err := r.redis.Set(r.context, uuid, description, 0)
	return err.Err()
}

func (r *ParticipantRepository) SetName(uuid int64, name string) error {
	query := fmt.Sprintf("UPDATE %s SET name=$1 WHERE uuid=$2", participantTable)
	_, err := r.db.Exec(query, name, uuid)

	return err
}

func (r *ParticipantRepository) SetPhoto(uuid int64, photo string) error {
	query := fmt.Sprintf("UPDATE %s SET photo=$1 WHERE uuid=$2", participantTable)
	_, err := r.db.Exec(query, photo, uuid)

	return err
}

func (r *ParticipantRepository) SetDescription(uuid int64, description string) error {
	query := fmt.Sprintf("UPDATE %s SET description=$1 WHERE uuid=$2", participantTable)
	_, err := r.db.Exec(query, description, uuid)

	return err
}

func (r *ParticipantRepository) GetCache(uuid string) (string, error) {
	value, err := r.redis.Get(r.context, uuid).Result()
	return value, err
}

func (r *ParticipantRepository) DeleteCache(uuid string) error {
	err := r.redis.Del(r.context, uuid)
	return err.Err()
}
