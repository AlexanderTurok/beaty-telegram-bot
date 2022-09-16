package service

import (
	"errors"
	"fmt"

	"github.com/AlexanderTurok/telegram-beaty-bot/internal/repository"
	telegram "github.com/AlexanderTurok/telegram-beaty-bot/pkg"
	"github.com/sirupsen/logrus"
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
		if err := s.repository.Create(uuid); err != nil {
			return err
		}
		if err := s.repository.Activate(uuid); err != nil {
			return err
		}
	}

	return nil
}

func (s *VoterService) GetParticipant(uuid int64) (telegram.Participant, error) {
	participant, err := s.repository.GetParticipant(uuid)
	if err != nil {
		logrus.Errorf("error while getting participant: %s", err)
		return telegram.Participant{}, errors.New("there is no availible profiles... Wait until someone adds new profile")
	}

	if err := participant.Validate(); err != nil {
		return telegram.Participant{}, err
	}

	if err := s.repository.DeleteParticipant(uuid, participant.Uuid); err != nil {
		return telegram.Participant{}, err
	}

	if err := s.repository.SetCache(fmt.Sprint(uuid), participant.Uuid); err != nil {
		return telegram.Participant{}, err
	}

	return participant, nil
}

func (s *VoterService) LikeParticipant(uuid int64) error {
	participantUuid, err := s.repository.GetCache(fmt.Sprint(uuid))
	if err != nil {
		return err
	}

	err = s.repository.LikeParticipant(participantUuid)

	return err
}
