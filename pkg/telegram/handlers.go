package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

const (
	start = "start"
)

func (b *Bot) handleCommands(message *tgbotapi.Message) error {
	switch message.Command() {
	case start:
		msg := tgbotapi.NewMessage(message.Chat.ID, "Hello, it's Beaty Bot. Choose your role to start:")
		_, err := b.bot.Send(msg)
		return err
	default:
		msg := tgbotapi.NewMessage(message.Chat.ID, "unknown command...")
		_, err := b.bot.Send(msg)
		return err
	}
}

// func (b *Bot) handleMessages(message *tgbotapi.Message) error {

// }

func (b *Bot) handleError(message *tgbotapi.Message, err error) {
	msg := tgbotapi.NewMessage(message.Chat.ID, err.Error())
	b.bot.Send(msg)
}
