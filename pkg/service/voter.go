package service

import (
	"strconv"

	"github.com/AlexanderTurok/telegram-beaty-bot"
	"github.com/AlexanderTurok/telegram-beaty-bot/pkg/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type VoterService struct {
	repository *repository.Repository
}

func NewVoterService(repository *repository.Repository) *VoterService {
	return &VoterService{
		repository: repository,
	}
}

func (v *VoterService) GetID(message *tgbotapi.Message) (int, error) {
	id, err := v.repository.Voter.GetCache(message.From.ID)
	if err != nil {
		return 0, err
	}

	if id == "" {
		if err := v.repository.Voter.SetCache(message.From.ID, "0"); err != nil {
			return 0, err
		}
	}

	intId, _ := strconv.Atoi(id)
	return intId, err
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
