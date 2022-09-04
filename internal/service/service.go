package service

import (
	"github.com/AlexanderTurok/telegram-beaty-bot"
	"github.com/AlexanderTurok/telegram-beaty-bot/pkg/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Participant interface {
	SetName(message *tgbotapi.Message) error
	SetPhoto(message *tgbotapi.Message) error
	SetDescription(message *tgbotapi.Message) error
	Create(message *tgbotapi.Message) error
	GetParticipant(uuid int) (*telegram.Participant, error)
	GetAllParticipants() (*[]telegram.Participant, error)
	UpdateParticipant(column, value string, uuid int) error
	DeleteParticipant(uuid int) error
	GetCache(uuid int) (string, error)
	SetCache(uuid int, value string) error
	DeleteCache(uuid int) error
}

type Voter interface {
	GetParticipant(uuid int) (*telegram.Participant, error)
}

type Service struct {
	Participant
	Voter
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Participant: NewParticipantService(repository),
		Voter:       NewVoterService(repository),
	}
}
