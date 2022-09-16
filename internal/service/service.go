package service

import (
	"github.com/AlexanderTurok/telegram-beaty-bot/internal/repository"
	telegram "github.com/AlexanderTurok/telegram-beaty-bot/pkg"
)

type Participant interface {
	Register(uuid int64) error
	Get(uuid int64) (telegram.Participant, error)
	Delete(uuid int64) error

	GetName(uuid int64, name string) error
	GetPhoto(uuid int64, photo string) error
	GetDescription(uuid int64, description string) error

	SetName(uuid int64, name string) error
	SetPhoto(uuid int64, photo string) error
	SetDescription(uuid int64, description string) error

	GetCache(uuid int64) (string, error)
	DeleteCache(uuid int64) error
}

type Voter interface {
	Create(uuid int64) error
	GetParticipant(uuid int64) (telegram.Participant, error)
	LikeParticipant(uuid int64) error
}

type Service struct {
	Participant
	Voter
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Participant: NewParticipantService(repository.Participant),
		Voter:       NewVoterService(repository.Voter),
	}
}
