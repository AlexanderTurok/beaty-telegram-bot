package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (b *Bot) handleCommands(message *tgbotapi.Message) error {
	switch message.Command() {
	case help:
		return b.handleHelp(message)
	case start:
		return b.handleStart(message)
	case support:
		return b.handleSupport(message)
	default:
		return b.handleDefaultCommand(message)
	}
}

func (b *Bot) handleHelp(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Available Commands: \n/start - to start bot \n/support - to report a problem")
	msg.ReplyMarkup = roleKeyboard
	_, err := b.bot.Send(msg)

	return err
}

func (b *Bot) handleStart(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Hello, it's Beaty Bot. Choose your role to start:")
	msg.ReplyMarkup = roleKeyboard
	_, err := b.bot.Send(msg)

	return err
}

func (b *Bot) handleSupport(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Contact @alexander_turok to resolve your problem.")
	msg.ReplyMarkup = roleKeyboard
	_, err := b.bot.Send(msg)

	return err
}

func (b *Bot) handleDefaultCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "unknown command...")
	_, err := b.bot.Send(msg)

	return err
}
