package bot

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (b *Bot) handleCommands(message *tgbotapi.Message) error {
	switch message.Command() {
	case start:
		err := b.handleStart(message)
		return err
	case support:
		err := b.handleSupport(message)
		return err
	default:
		err := b.handleDefault(message)
		return err
	}
}

func (b *Bot) handleStart(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Hello, it's Beaty Bot. Choose your role to start:")
	msg.ReplyMarkup = roleKeyboard
	_, err := b.bot.Send(msg)

	return fmt.Errorf("error in start handler: %s", err)
}

func (b *Bot) handleSupport(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Contact @alexander_turok to resolve your problem.")
	msg.ReplyMarkup = roleKeyboard
	_, err := b.bot.Send(msg)

	return fmt.Errorf("error in support handler: %s", err)
}

func (b *Bot) handleDefault(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "unknown command...")
	_, err := b.bot.Send(msg)

	return fmt.Errorf("error in default-command handler: %s", err)
}
