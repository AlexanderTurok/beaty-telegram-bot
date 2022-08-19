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

func (p *ParticipantDBService) GetParticipant(uuid int) (*telegram.Participant, error) {
	return nil, nil
}

func (p *ParticipantDBService) AddParticipant(uuid int) error {
	return nil
}

func (p *ParticipantDBService) UpdateParticipant(column, value string, uuid int) error {
	return nil
}

func (p *ParticipantDBService) DeleteParticipant(uuid int) error {
	return nil
}
