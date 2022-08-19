package service

import (
	"github.com/AlexanderTurok/telegram-beaty-bot"
	"github.com/AlexanderTurok/telegram-beaty-bot/pkg/repository"
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
	UpdateParticipant(uuid string) error
}

type VoterCache interface {
	GetCache(uuid string) (string, error)
	SetCache(uuid, value string) error
	DeleteCache(uuid string) error
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
