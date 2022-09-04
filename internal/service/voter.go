package service

import (
	"github.com/AlexanderTurok/telegram-beaty-bot/internal/repository"
	telegram "github.com/AlexanderTurok/telegram-beaty-bot/pkg"
)

type VoterService struct {
	repository repository.Voter
}

func NewVoterService(repository repository.Voter) *VoterService {
	return &VoterService{
		repository: repository,
	}
}

func (s *VoterService) Create(uuid int64) error {
	exists, err := s.repository.IsExists(uuid)
	if err != nil {
		return err
	}
	if !exists {
		return s.repository.Create(uuid)
	}

	return nil
}

func (s *VoterService) GetParticipant(uuid int64) (telegram.Participant, error) {
	participant, err := s.repository.GetParticipant(uuid)
	if err != nil {
		return telegram.Participant{}, err
	}

	if err := participant.Validate(); err != nil {
		return telegram.Participant{}, err
	}

	return participant, err
}
