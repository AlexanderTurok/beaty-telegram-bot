package bot

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (b *Bot) handleCache(message *tgbotapi.Message, value string) error {
	switch value {
	case "name":
		err := b.handleName(message)
		return err
	case "photo":
		err := b.handlePhoto(message)
		return err
	case "description":
		err := b.handleDescription(message)
		return err
	default:
		err := b.handleDefaultCache(message)
		return err
	}
}

func (b *Bot) handleName(message *tgbotapi.Message) error {
	if err := b.service.Participant.SetName(message); err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, "Your name successfully updated!")
	msg.ReplyMarkup = registrationKeyboard
	_, err := b.bot.Send(msg)

	return fmt.Errorf("error in name handler: %s", err)
}

func (b *Bot) handlePhoto(message *tgbotapi.Message) error {
	if err := b.service.Participant.SetPhoto(message); err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, "Your photo successfully updated!")
	msg.ReplyMarkup = registrationKeyboard
	_, err := b.bot.Send(msg)

	return fmt.Errorf("error in photo handler: %s", err)
}

func (b *Bot) handleDescription(message *tgbotapi.Message) error {
	if err := b.service.Participant.SetDescription(message); err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, "Your description successfully updated!")
	msg.ReplyMarkup = registrationKeyboard
	_, err := b.bot.Send(msg)

	return fmt.Errorf("error in description handler: %s", err)
}

func (b *Bot) handleDefaultCache(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Error in cache handler. Contact with @alexander_turok.")
	_, err := b.bot.Send(msg)

	return fmt.Errorf("error in default-cache handler: %s", err)
}
