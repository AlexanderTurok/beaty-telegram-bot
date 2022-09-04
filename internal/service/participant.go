package service

import (
	"github.com/AlexanderTurok/telegram-beaty-bot/internal/repository"
	telegram "github.com/AlexanderTurok/telegram-beaty-bot/pkg"
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
	return err
}

func (p *ParticipantService) SetPhoto(message *tgbotapi.Message) error {
	if err := p.repository.Participant.DeleteCache(message.From.ID); err != nil {
		return err
	}

	fileID := (*message.Photo)[0].FileID
	err := p.repository.Participant.UpdateParticipant("photo", fileID, message.From.ID)

	return err
}

func (p *ParticipantService) SetDescription(message *tgbotapi.Message) error {
	if err := p.repository.Participant.DeleteCache(message.From.ID); err != nil {
		return err
	}

	err := p.repository.Participant.UpdateParticipant("description", message.Text, message.From.ID)
	return err
}

func (p *ParticipantService) Create(message *tgbotapi.Message) error {
	exists, err := p.repository.Participant.IsExists(message.From.ID)
	if !exists {
		if err := p.repository.Participant.Create(message.From.ID); err != nil {
			return err
		}
	}

	return err
}

func (p *ParticipantService) IsExists(uuid int) (bool, error) {
	exists, err := p.repository.Participant.IsExists(uuid)
	return exists, err
}

func (p *ParticipantService) GetParticipant(uuid int) (telegram.Participant, error) {
	participant, err := p.repository.Participant.GetParticipant(uuid)
	return participant, err
}

func (v *ParticipantService) GetAllParticipants() ([]telegram.Participant, error) {
	participants, err := v.repository.Participant.GetAllParticipants()
	return participants, err
}

func (p *ParticipantService) UpdateParticipant(column, value string, uuid int) error {
	err := p.repository.Participant.UpdateParticipant(column, value, uuid)
	return err
}

func (p *ParticipantService) DeleteParticipant(uuid int) error {
	err := p.repository.Participant.DeleteParticipant(uuid)
	return err
}

func (p *ParticipantService) SetCache(uuid int, value string) error {
	err := p.repository.Participant.SetCache(uuid, value)
	return err
}

func (p *ParticipantService) GetCache(uuid int) (string, error) {
	cache, err := p.repository.Participant.GetCache(uuid)
	return cache, err
}

func (p *ParticipantService) DeleteCache(uuid int) error {
	err := p.repository.Participant.DeleteCache(uuid)
	return err
}
