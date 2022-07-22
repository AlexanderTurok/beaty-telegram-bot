package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

const (
	start = "start"
)

func (b *Bot) handleCommand(message *tgbotapi.Message) {
	switch message.Command() {
	case start:
		msg := tgbotapi.NewMessage(message.Chat.ID, "Hello, it's Beaty Bot. Choose your role to start:")
		b.bot.Send(msg)

	default:
		msg := tgbotapi.NewMessage(message.Chat.ID, "unknown command...")
		b.bot.Send(msg)
	}
}
