package service

import (
	"github.com/AlexanderTurok/telegram-beaty-bot"
	"github.com/AlexanderTurok/telegram-beaty-bot/pkg/repository"
)

type ParticipantService struct {
	repository *repository.Repository
}

func NewParticipantService(repository *repository.Repository) *ParticipantService {
	return &ParticipantService{
		repository: repository,
	}
}

func (p *ParticipantService) IsParticipant(uuid int) (bool, error) {
	return false, nil
}

func (p *ParticipantService) GetParticipant(uuid int) (*telegram.Participant, error) {
	return nil, nil
}

func (v *ParticipantService) GetAllParticipants() (*[]telegram.Participant, error) {
	return nil, nil
}

func (p *ParticipantService) AddParticipant(uuid int) error {
	return nil
}

func (p *ParticipantService) UpdateParticipant(column, value string, uuid int) error {
	return nil
}

func (p *ParticipantService) DeleteParticipant(uuid int) error {
	return nil
}

func (p *ParticipantService) SetCache(uuid int, value string) error {
	return nil
}

func (p *ParticipantService) GetCache(uuid int) (string, error) {
	return "", nil
}

func (p *ParticipantService) DeleteCache(uuid int) error {
	return nil
}
