package bot

import (
	"github.com/AlexanderTurok/telegram-beaty-bot/internal/service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
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
		logrus.Fatalf("error in updates channel: %s", err)
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			err := b.handleCommands(update.Message)
			if err != nil {
				logrus.Errorf("error in command handler: %s", err)
			}
			continue
		}

		// check if participants message has cache
		if value, _ := b.service.Participant.GetCache(update.Message.Chat.ID); value != name && value != photo && value != description {
			if err := b.handleCache(update.Message, value); err != nil {
				logrus.Errorf("error in cache handler: %s", err)
			}
			continue
		}

		if err := b.handleMessages(update.Message); err != nil {
			logrus.Errorf("error in message handler: %s", err)
		}
	}
}
