package repository

import (
	"context"
	"database/sql"

	"github.com/AlexanderTurok/telegram-beaty-bot"
	"github.com/go-redis/redis/v9"
)

type ParticipantData interface {
	GetParticipant(uuid string) (*telegram.Participant, error)
	AddParticipant(uuid string) error
	UpdateParticipant(column, value, uuid string) error
	DeleteParticipant(uuid string) error
}

type ParticipantCache interface {
	GetCache(uuid string) (string, error)
	SetCache(uuid, value string) error
	DeleteCache(uuid string) error
}

type VoterData interface {
	GetParticipant(uuid string) (*telegram.Participant, error)
	UpdateParticipant(column, value, uuid string) error
}

type VoterCache interface {
	GetCache(uuid string) (string, error)
	SetCache(uuid, value string) error
	DeleteCache(uuid string) error
}

type Repository struct {
	ParticipantData
	ParticipantCache
	VoterData
	VoterCache
}

func NewRepository(context context.Context, db *sql.DB, redis *redis.Client) *Repository {
	return &Repository{
		ParticipantData:  NewParticipantDB(db),
		ParticipantCache: NewParticipantRD(context, redis),
		VoterData:        NewVoterDB(db),
		VoterCache:       NewVoterRD(context, redis),
	}
}
