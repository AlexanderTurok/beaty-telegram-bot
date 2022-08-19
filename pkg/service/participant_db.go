package service

import (
	"github.com/AlexanderTurok/telegram-beaty-bot"
	"github.com/AlexanderTurok/telegram-beaty-bot/pkg/repository"
)

type ParticipantDBService struct {
	repository *repository.Repository
}

func NewParticipantDBService(repository *repository.Repository) *ParticipantDBService {
	return &ParticipantDBService{
		repository: repository,
	}
}

func (p *ParticipantDBService) GetParticipant(uuid string) (*telegram.Participant, error) {
	return nil, nil
}

func (p *ParticipantDBService) AddParticipant(uuid string) error {
	return nil
}

func (p *ParticipantDBService) UpdateParticipant(column, value, uuid string) error {
	return nil
}

func (p *ParticipantDBService) DeleteParticipant(uuid string) error {
	return nil
}
