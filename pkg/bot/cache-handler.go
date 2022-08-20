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
		if err := b.service.ParticipantCache.DeleteCache(message.From.ID); err != nil {
			return err
		}

		err := b.service.ParticipantData.UpdateParticipant("photo", (*message.Photo)[0].FileID, message.From.ID)
		if err != nil {
			return err
		} else {
			msg := tgbotapi.NewMessage(message.Chat.ID, "Your photo successfully updated!")
			msg.ReplyMarkup = registrationKeyboard
			_, err := b.bot.Send(msg)
			return err
		}
	case "description":
		if err := b.service.ParticipantCache.DeleteCache(message.From.ID); err != nil {
			return err
		}

		err := b.service.ParticipantData.UpdateParticipant("information", message.Text, message.From.ID)
		if err != nil {
			return err
		} else {
			msg := tgbotapi.NewMessage(message.Chat.ID, "Your description successfully updated!")
			msg.ReplyMarkup = registrationKeyboard
			_, err := b.bot.Send(msg)
			return err
		}
	default:
		msg := tgbotapi.NewMessage(message.Chat.ID, "something went from with cache handler")
		_, err := b.bot.Send(msg)
		return err
	}
}

func (b *Bot) handleName(message *tgbotapi.Message) error {
	if err := b.service.Participant.SetParticipantName(message); err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, "Your name successfully updated!")
	msg.ReplyMarkup = registrationKeyboard
	_, err := b.bot.Send(msg)

	return fmt.Errorf("error in name handler: %s", err)
}
