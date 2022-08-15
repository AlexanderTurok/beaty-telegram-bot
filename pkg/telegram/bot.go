package telegram

import (
	"context"
	"database/sql"
	"log"

	"github.com/go-redis/redis/v9"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Bot struct {
	bot      *tgbotapi.BotAPI
	redis    *redis.Client
	postgres *sql.DB
	ctx      context.Context
}

func NewBot(bot *tgbotapi.BotAPI, redis *redis.Client, postgres *sql.DB, ctx context.Context) *Bot {
	return &Bot{
		bot:      bot,
		redis:    redis,
		postgres: postgres,
		ctx:      ctx,
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

		if value := b.getCache(update.Message.From.ID); value != "" {
			err := b.handleCache(update.Message, value)
			if err != nil {
				log.Fatalf("error in cache handler: %s", err)
			}
			continue
		}

		err := b.handleMessages(update.Message)
		if err != nil {
			log.Fatalf("error in message handler: %s", err.Error())
		}
	}
}
