package repository

import (
	"context"
	"database/sql"

	"github.com/AlexanderTurok/telegram-beaty-bot"
	"github.com/go-redis/redis/v9"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Participant interface {
	SetParticipantName(message *tgbotapi.Message) error

	IsParticipant(uuid int) (bool, error)
	GetParticipant(uuid int) (*telegram.Participant, error)
	GetAllParticipants() (*[]telegram.Participant, error)
	AddParticipant(uuid int) error
	UpdateParticipant(column, value string, uuid int) error
	DeleteParticipant(uuid int) error
	GetCache(uuid int) (string, error)
	SetCache(uuid int, value string) error
	DeleteCache(uuid int) error
}

type Voter interface {
	GetParticipant(uuid int) (*telegram.Participant, error)
	UpdateParticipant(column, value string, uuid int) error
	GetCache(uuid int) (string, error)
	SetCache(uuid int, value string) error
	DeleteCache(uuid int) error
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
