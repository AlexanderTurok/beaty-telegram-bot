package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/AlexanderTurok/telegram-beaty-bot"
	"github.com/go-redis/redis/v9"
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

func (p *ParticipantRepository) IsParticipant(uuid int) (bool, error) {
	var exists bool
	query := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM participants WHERE uuid='%d');", uuid)
	err := p.db.QueryRow(query).Scan(&exists)

	return exists, err
}

func (p *ParticipantRepository) GetParticipant(uuid int) (*telegram.Participant, error) {
	var par telegram.Participant
	query := fmt.Sprintf("SELECT * FROM participants WHERE uuid=%d;", uuid)
	err := p.db.QueryRow(query).Scan(&par.Id, &par.Uuid, &par.Nickname, &par.Photo, &par.Information, &par.Votes)

	return &par, err
}

func (p *ParticipantRepository) GetAllParticipants() (*[]telegram.Participant, error) {
	query := "SELECT * FROM participants;"
	rows, err := p.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var participants []telegram.Participant

	for rows.Next() {
		var par telegram.Participant
		if err := rows.Scan(&par.Id, &par.Uuid, &par.Nickname,
			&par.Photo, &par.Information, &par.Votes); err != nil {
			return &participants, err
		}
		participants = append(participants, par)
	}

	if err = rows.Err(); err != nil {
		return &participants, err
	}

	return &participants, nil
}

func (p *ParticipantRepository) AddParticipant(uuid int) error {
	query := fmt.Sprintf("INSERT INTO participants (uuid) VALUES (%d);", uuid)
	_, err := p.db.Exec(query)

	return err
}

func (p *ParticipantRepository) UpdateParticipant(column, value string, uuid int) error {
	query := fmt.Sprintf("UPDATE participants SET %s='%s' WHERE uuid=%d", column, value, uuid)
	_, err := p.db.Exec(query)

	return err
}

func (p *ParticipantRepository) DeleteParticipant(uuid int) error {
	query := fmt.Sprintf("DELETE FROM participants WHERE uuid=%d;", uuid)
	_, err := p.db.Exec(query)

	return err
}

func (p *ParticipantRepository) SetCache(uuid int, value string) error {
	err := p.redis.Set(p.context, fmt.Sprint(uuid), value, 0)
	return err.Err()
}

func (p *ParticipantRepository) GetCache(uuid int) (string, error) {
	value, err := p.redis.Get(p.context, fmt.Sprint(uuid)).Result()
	return value, err
}

func (p *ParticipantRepository) DeleteCache(uuid int) error {
	err := p.redis.Del(p.context, fmt.Sprint(uuid))
	return err.Err()
}
