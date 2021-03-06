package telegram

import (
	"github.com/AlexanderTurok/beaty-telegram-bot/pkg/data"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	start = "start"
	miss  = "New Miss!"
	voter = "Voter!"
)

var roleKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(miss),
		tgbotapi.NewKeyboardButton(voter),
	),
)

func (b *Bot) handleCommands(message *tgbotapi.Message) error {
	switch message.Command() {
	case start:
		msg := tgbotapi.NewMessage(message.Chat.ID, "Hello, it's Beaty Bot. Choose your role to start:")
		msg.ReplyMarkup = roleKeyboard
		_, err := b.bot.Send(msg)
		return err
	default:
		msg := tgbotapi.NewMessage(message.Chat.ID, "unknown command...")
		_, err := b.bot.Send(msg)
		return err
	}
}

func (b *Bot) handleMessages(message *tgbotapi.Message) error {
	switch message.Text {
	case miss:
		exists, err := data.IsParticipantRowExists("SELECT uuid FROM participant", message.Chat.ID)
		if exists {
			return err
		}
		return err
	case voter:
		return nil
	default:
		msg := tgbotapi.NewMessage(message.Chat.ID, "unknown message...")
		_, err := b.bot.Send(msg)
		return err
	}
}

func (b *Bot) handleError(message *tgbotapi.Message, err error) {
	msg := tgbotapi.NewMessage(message.Chat.ID, err.Error())
	b.bot.Send(msg)
}
