package service

import (
	"github.com/AlexanderTurok/telegram-beaty-bot"
	"github.com/AlexanderTurok/telegram-beaty-bot/pkg/repository"
)

type ParticipantData interface {
	GetParticipant(uuid int) (*telegram.Participant, error)
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
	UpdateParticipant(uuid int) error
}

type VoterCache interface {
	GetCache(uuid int) (string, error)
	SetCache(uuid int, value string) error
	DeleteCache(uuid int) error
}

type Service struct {
	ParticipantData
	ParticipantCache
	VoterData
	VoterCache
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		ParticipantData:  NewParticipantDBService(repository),
		ParticipantCache: NewParticipantRDService(repository),
		VoterData:        NewVoterDBService(repository),
		VoterCache:       NewVoterRDService(repository),
	}
}
