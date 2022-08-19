package service

import (
	"github.com/AlexanderTurok/telegram-beaty-bot/pkg/repository"
)

type VoterRDService struct {
	repository *repository.Repository
}

func NewVoterRDService(repository *repository.Repository) *VoterRDService {
	return &VoterRDService{
		repository: repository,
	}
}

func (v *VoterRDService) SetCache(uuid int, value string) error {
	return nil
}

func (v *VoterRDService) GetCache(uuid int) (string, error) {
	return "", nil
}

func (v *VoterRDService) DeleteCache(uuid int) error {
	return nil
}
