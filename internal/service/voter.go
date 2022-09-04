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
	if err := v.create(uuid); err != nil {
		return nil, err
	}

	participantUUID, err := v.repository.Voter.GetParticipantUUID(uuid)
	if err != nil {
		return nil, err
	}

	participant, err := v.repository.GetParticipant(participantUUID[0])
	if err != nil {
		return participant, err
	}

	err = v.repository.Voter.DeleteParticipant(uuid, participantUUID[0])

	return participant, err
}

func (v *VoterService) create(uuid int) error {
	exists, err := v.repository.Voter.IsExists(uuid)
	if !exists {
		if err := v.repository.Voter.Create(uuid); err != nil {
			return err
		}
	}

	return err
}
