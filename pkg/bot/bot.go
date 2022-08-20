package bot

import (
	"log"

	"github.com/AlexanderTurok/telegram-beaty-bot/pkg/service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Bot struct {
	bot     *tgbotapi.BotAPI
	service *service.Service
}

func NewBot(bot *tgbotapi.BotAPI, service *service.Service) *Bot {
	return &Bot{
		bot:     bot,
		service: service,
	}
}

func (b *Bot) Start() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := b.bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatalf("error in updates channel: %s", err.Error())
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			err := b.handleCommands(update.Message)
			if err != nil {
				log.Fatalf("error in command handler: %s", err.Error())
			}
			continue
		}

		if value, _ := b.service.Participant.GetCache(update.Message.From.ID); value != "" {
			if err := b.handleCache(update.Message, value); err != nil {
				log.Fatalf("error in cache handler: %s", err)
			}
			continue
		}

		if err := b.handleMessages(update.Message); err != nil {
			log.Fatalf("error in message handler: %s", err.Error())
		}
	}
}
