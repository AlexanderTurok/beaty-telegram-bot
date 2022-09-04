package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (b *Bot) handleCache(message *tgbotapi.Message, value string) error {
	switch value {
	case nameCache:
		err := b.setName(message)
		return err
	case photoCache:
		err := b.setPhoto(message)
		return err
	case descriptionCache:
		err := b.setDescription(message)
		return err
	default:
		err := b.handleDefaultCache(message)
		return err
	}
}

func (b *Bot) setName(message *tgbotapi.Message) error {
	if err := b.service.Participant.SetName(message.Chat.ID, message.Text); err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, "Your name successfully updated!")
	msg.ReplyMarkup = registrationKeyboard
	_, err := b.bot.Send(msg)

	return err
}

func (b *Bot) setPhoto(message *tgbotapi.Message) error {
	photoId := (*message.Photo)[0].FileID
	if err := b.service.Participant.SetPhoto(message.Chat.ID, photoId); err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, "Your photo successfully updated!")
	msg.ReplyMarkup = registrationKeyboard
	_, err := b.bot.Send(msg)

	return err
}

func (b *Bot) setDescription(message *tgbotapi.Message) error {
	if err := b.service.Participant.SetDescription(message.Chat.ID, message.Text); err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, "Your description successfully updated!")
	msg.ReplyMarkup = registrationKeyboard
	_, err := b.bot.Send(msg)

	return err
}

func (b *Bot) handleDefaultCache(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Error in cache handler. Contact with @alexander_turok.")
	_, err := b.bot.Send(msg)

	return err
}
