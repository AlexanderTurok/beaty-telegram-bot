package service

import (
	"github.com/AlexanderTurok/telegram-beaty-bot/internal/repository"
)

type VoterService struct {
	repository repository.Voter
}

func NewVoterService(repository repository.Voter) *VoterService {
	return &VoterService{
		repository: repository,
	}
}
