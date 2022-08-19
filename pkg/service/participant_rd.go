package service

import (
	"github.com/AlexanderTurok/telegram-beaty-bot/pkg/repository"
)

type ParticipantRDService struct {
	repository *repository.Repository
}

func NewParticipantRDService(repository *repository.Repository) *ParticipantRDService {
	return &ParticipantRDService{
		repository: repository,
	}
}

func (p *ParticipantRDService) SetCache(uuid int, value string) error {
	return nil
}

func (p *ParticipantRDService) GetCache(uuid int) (string, error) {
	return "", nil
}

func (p *ParticipantRDService) DeleteCache(uuid int) error {
	return nil
}
