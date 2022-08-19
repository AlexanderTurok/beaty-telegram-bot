package repository

import (
	"context"
	"database/sql"

	"github.com/AlexanderTurok/telegram-beaty-bot"
	"github.com/go-redis/redis/v9"
)

type ParticipantData interface {
	IsParticipant(uuid int) (bool, error)
	GetParticipant(uuid int) (*telegram.Participant, error)
	GetAllParticipants() (*[]telegram.Participant, error)
	AddParticipant(uuid int) error
	UpdateParticipant(column, value string, uuid int) error
	DeleteParticipant(uuid int) error
}

type ParticipantCache interface {
	GetCache(uuid int) (string, error)
	SetCache(uuid int, value string) error
	DeleteCache(uuid int) error
}

type VoterData interface {
	GetParticipant(uuid int) (*telegram.Participant, error)
	UpdateParticipant(column, value string, uuid int) error
}

type VoterCache interface {
	GetCache(uuid int) (string, error)
	SetCache(uuid int, value string) error
	DeleteCache(uuid int) error
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
