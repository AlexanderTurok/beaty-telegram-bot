package repository

import (
	"context"
	"database/sql"

	telegram "github.com/AlexanderTurok/telegram-beaty-bot/pkg"
	"github.com/go-redis/redis/v9"
)

type Participant interface {
	IsExists(uuid int) (bool, error)
	GetParticipant(uuid int) (telegram.Participant, error)
	GetAllParticipants() ([]telegram.Participant, error)
	Create(uuid int) error
	UpdateParticipant(column, value string, uuid int) error
	DeleteParticipant(uuid int) error
	GetCache(uuid int) (string, error)
	SetCache(uuid int, value string) error
	DeleteCache(uuid int) error
}

type Voter interface {
	Create(uuid int) error
	GetParticipantUUID(uuid int) ([]int, error)
	DeleteParticipant(voterUuid, ParticipantUuid int) error
	IsExists(uuid int) (bool, error)
}

type Repository struct {
	Participant
	Voter
}

func NewRepository(context context.Context, db *sql.DB, redis *redis.Client) *Repository {
	return &Repository{
		Participant: NewParticipantRepository(context, db, redis),
		Voter:       NewVoterRepository(context, db, redis),
	}
}
