package service

import (
	"fmt"

	"github.com/AlexanderTurok/telegram-beaty-bot"
	"github.com/AlexanderTurok/telegram-beaty-bot/pkg/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type ParticipantService struct {
	repository *repository.Repository
}

func NewParticipantService(repository *repository.Repository) *ParticipantService {
	return &ParticipantService{
		repository: repository,
	}
}

func (p *ParticipantService) SetName(message *tgbotapi.Message) error {
	if err := p.repository.Participant.DeleteCache(message.From.ID); err != nil {
		return err
	}

	err := p.repository.Participant.UpdateParticipant("name", message.Text, message.From.ID)
	return fmt.Errorf("error while setting name in services: %s", err)
}

func (p *ParticipantService) SetPhoto(message *tgbotapi.Message) error {
	if err := p.repository.Participant.DeleteCache(message.From.ID); err != nil {
		return err
	}

	fileID := (*message.Photo)[0].FileID
	err := p.repository.Participant.UpdateParticipant("photo", fileID, message.From.ID)

	return fmt.Errorf("error while setting photo in services: %s", err)
}

func (p *ParticipantService) SetDescription(message *tgbotapi.Message) error {
	if err := p.repository.Participant.DeleteCache(message.From.ID); err != nil {
		return err
	}

	err := p.repository.Participant.UpdateParticipant("description", message.Text, message.From.ID)
	return fmt.Errorf("error while setting description in services: %s", err)
}

func (p *ParticipantService) Register(message *tgbotapi.Message) error {
	exists, err := p.repository.Participant.IsParticipant(message.From.ID)
	if !exists {
		if err := p.repository.Participant.AddParticipant(message.From.ID); err != nil {
			return err
		}
	}

	return fmt.Errorf("error while register in services: %s", err)
}

func (p *ParticipantService) IsParticipant(uuid int) (bool, error) {
	exists, err := p.repository.IsParticipant(uuid)
	return exists, fmt.Errorf("error while getting exist of participant in services: %s", err)
}

func (p *ParticipantService) GetParticipant(uuid int) (*telegram.Participant, error) {
	participant, err := p.repository.Participant.GetParticipant(uuid)
	return participant, fmt.Errorf("error while getting participant in services: %s", err)
}

func (v *ParticipantService) GetAllParticipants() (*[]telegram.Participant, error) {
	participants, err := v.repository.Participant.GetAllParticipants()
	return participants, fmt.Errorf("error while getting all participant in services: %s", err)
}

func (p *ParticipantService) UpdateParticipant(column, value string, uuid int) error {
	err := p.repository.Participant.UpdateParticipant(column, value, uuid)
	return fmt.Errorf("error while updating participant in services: %s", err)
}

func (p *ParticipantService) DeleteParticipant(uuid int) error {
	err := p.repository.Participant.DeleteParticipant(uuid)
	return fmt.Errorf("error while deleting participant in services: %s", err)
}

func (p *ParticipantService) SetCache(uuid int, value string) error {
	err := p.repository.Participant.SetCache(uuid, value)
	return fmt.Errorf("error while set cache in services: %s", err)
}

func (p *ParticipantService) GetCache(uuid int) (string, error) {
	cache, err := p.repository.Participant.GetCache(uuid)
	return cache, fmt.Errorf("error while set cache in services: %s", err)
}

func (p *ParticipantService) DeleteCache(uuid int) error {
	err := p.repository.Participant.DeleteCache(uuid)
	return fmt.Errorf("error while set cache in services: %s", err)
}
