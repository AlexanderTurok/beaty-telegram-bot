package repository

import (
	"context"

	telegram "github.com/AlexanderTurok/telegram-beaty-bot/pkg"
	"github.com/go-redis/redis/v9"
	"github.com/jmoiron/sqlx"
)

type Participant interface {
	Register(uuid int64) error
	IsExists(uuid int64) (bool, error)
	Activate(uuid int64) error

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
	Create(uuid int64) error
	Activate(uuid int64) error
	IsExists(uuid int64) (bool, error)

	GetParticipant(uuid int64) (telegram.Participant, error)
	LikeParticipant(uuid string) error
	DeleteParticipant(voterUuid int64, participantUuid string) error

	SetCache(voterUuid, participantUuid string) error
	GetCache(uuid string) (string, error)
}

type Repository struct {
	Participant
	Voter
}

func NewRepository(context context.Context, db *sqlx.DB, redis *redis.Client) *Repository {
	return &Repository{
		Participant: NewParticipantRepository(context, db, redis),
		Voter:       NewVoterRepository(context, db, redis),
	}
}
