package main

import (
	"log"

	"github.com/AlexanderTurok/beaty-telegram-bot/pkg/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	botApi, err := tgbotapi.NewBotAPI("MyAwesomeBotToken")
	if err != nil {
		log.Panic(err)
	}

	botApi.Debug = true

	bot := telegram.NewBot(botApi)

	bot.Start()
}
