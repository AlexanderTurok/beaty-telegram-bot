package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/AlexanderTurok/telegram-beaty-bot"
	"github.com/go-redis/redis/v9"
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

func (v *VoterRepository) GetParticipant(uuid int) (*telegram.Participant, error) {
	var p telegram.Participant
	query := fmt.Sprintf("SELECT * FROM participants WHERE uuid=%d;", uuid)
	err := v.db.QueryRow(query).Scan(&p.Id, &p.Uuid, &p.Nickname, &p.Photo, &p.Information, &p.Votes)

	return &p, err
}

func (v *VoterRepository) UpdateParticipant(column, value string, uuid int) error {
	query := fmt.Sprintf("UPDATE participants SET %s='%s' WHERE uuid=%d", column, value, uuid)
	_, err := v.db.Exec(query)

	return err
}

func (v *VoterRepository) SetCache(uuid int, value string) error {
	err := v.redis.Set(v.context, fmt.Sprint(uuid), value, 0)
	return err.Err()
}

func (v *VoterRepository) GetCache(uuid int) (string, error) {
	value, err := v.redis.Get(v.context, fmt.Sprint(uuid)).Result()
	return value, err
}

func (v *VoterRepository) DeleteCache(uuid int) error {
	err := v.redis.Del(v.context, fmt.Sprint(uuid))
	return err.Err()
}
