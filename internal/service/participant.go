package service

import (
	"fmt"

	"github.com/AlexanderTurok/telegram-beaty-bot/internal/repository"
	telegram "github.com/AlexanderTurok/telegram-beaty-bot/pkg"
)

type ParticipantService struct {
	repository repository.Participant
}

func NewParticipantService(repository repository.Participant) *ParticipantService {
	return &ParticipantService{
		repository: repository,
	}
}

func (s *ParticipantService) Register(uuid int64) error {
	exists, err := s.repository.IsExists(uuid)
	if err != nil {
		return err
	}
	if !exists {
		return s.repository.Register(uuid)
	}

	return nil
}

func (s *ParticipantService) Get(uuid int64) (telegram.Participant, error) {
	return s.repository.Get(uuid)
}

func (s *ParticipantService) Delete(uuid int64) error {
	return s.repository.Delete(uuid)
}

func (s *ParticipantService) GetName(uuid int64, name string) error {
	return s.repository.GetName(fmt.Sprint(uuid), name)
}

func (s *ParticipantService) GetPhoto(uuid int64, photo string) error {
	return s.repository.GetPhoto(fmt.Sprint(uuid), photo)
}

func (s *ParticipantService) GetDescription(uuid int64, description string) error {
	return s.repository.GetDescription(fmt.Sprint(uuid), description)
}

func (s *ParticipantService) SetName(uuid int64, name string) error {
	return s.repository.SetName(uuid, name)
}

func (s *ParticipantService) SetPhoto(uuid int64, photo string) error {
	return s.repository.SetPhoto(uuid, photo)
}

func (s *ParticipantService) SetDescription(uuid int64, description string) error {
	return s.repository.SetDescription(uuid, description)
}

func (s *ParticipantService) GetCache(uuid int64) (string, error) {
	return s.repository.GetCache(fmt.Sprint(uuid))
}

func (s *ParticipantService) DeleteCache(uuid int64) error {
	return s.repository.DeleteCache(fmt.Sprint(uuid))
}