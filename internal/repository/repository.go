package repository

import (
	"context"
	"database/sql"

	telegram "github.com/AlexanderTurok/telegram-beaty-bot/pkg"
	"github.com/go-redis/redis/v9"
)

type Participant interface {
	Register(uuid int64) error
	IsExists(uuid int64) (bool, error)
	Get(uuid int64) (telegram.Participant, error)
	Delete(uuid int64) error

	GetName(uuid, name string) error
	GetPhoto(uuid, photo string) error
	GetDescription(uuid, description string) error

	SetName(uuid int64, name string) error
	SetPhoto(uuid int64, photo string) error
	SetDescription(uuid int64, description string) error

	GetCache(uuid string) (string, error)
	DeleteCache(uuid string) error
}

type Voter interface {
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