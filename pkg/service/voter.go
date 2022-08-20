package service

import (
	"github.com/AlexanderTurok/telegram-beaty-bot"
	"github.com/AlexanderTurok/telegram-beaty-bot/pkg/repository"
)

type VoterService struct {
	repository *repository.Repository
}

func NewVoterService(repository *repository.Repository) *VoterService {
	return &VoterService{
		repository: repository,
	}
}

func (v *VoterService) GetParticipant(uuid int) (*telegram.Participant, error) {
	return nil, nil
}

func (v *VoterService) UpdateParticipant(uuid int) error {
	return nil
}

func (v *VoterService) SetCache(uuid int, value string) error {
	return nil
}

func (v *VoterService) GetCache(uuid int) (string, error) {
	return "", nil
}

func (v *VoterService) DeleteCache(uuid int) error {
	return nil
}
