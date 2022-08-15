package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	// command list
	start = "start"

	// choose role
	miss  = "New Miss!"
	voter = "Voter!"

	// register
	name        = "Change a Name"
	photo       = "Add a Photo"
	description = "Give a Description"
	back        = "Go Back"
)

var roleKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(miss),
		tgbotapi.NewKeyboardButton(voter),
	),
)

var registrationKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(name),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(photo),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(description),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(back),
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
		exists, err := b.isParticipantInDB(message.From.ID)
		if err != nil {
			return err
		}
		if !exists {
			err := b.addParticipantToDB("uuid", message.From.ID)
			return err
		} else {
			msg := tgbotapi.NewMessage(message.Chat.ID, "Here is available methods:")
			msg.ReplyMarkup = registrationKeyboard
			_, err := b.bot.Send(msg)
			return err
		}
	case name:
		msg := tgbotapi.NewMessage(message.Chat.ID, "Send your name visible to others")
		_, err := b.bot.Send(msg)
		if err != nil {
			return err
		}

		err = b.setCache(message.From.ID, "name")
		return err
	case photo:
		msg := tgbotapi.NewMessage(message.Chat.ID, "Waiting for you photo...")
		_, err := b.bot.Send(msg)
		if err != nil {
			return err
		}

		err = b.setCache(message.From.ID, "photo")
		return err
	case description:
		msg := tgbotapi.NewMessage(message.Chat.ID, "Describe yourself")
		_, err := b.bot.Send(msg)
		if err != nil {
			return err
		}

		err = b.setCache(message.From.ID, "description")
		return err
	case voter:
		return nil
	default:
		msg := tgbotapi.NewMessage(message.Chat.ID, "unknown message...")
		_, err := b.bot.Send(msg)
		return err
	}
}

func (b *Bot) handleCache(message *tgbotapi.Message, value string) error {
	switch value {
	case "name":
		if err := b.deleteCache(message.From.ID); err != nil {
			return err
		}

		err := b.updateParticipantInDB("nickname", message.Text, message.From.ID)
		if err != nil {
			return err
		} else {
			msg := tgbotapi.NewMessage(message.Chat.ID, "Your name successfully updated!")
			msg.ReplyMarkup = registrationKeyboard
			_, err := b.bot.Send(msg)
			return err
		}
	case "photo":
		if err := b.deleteCache(message.From.ID); err != nil {
			return err
		}
		imageUrl, err := b.bot.GetFileDirectURL((*message.Photo)[0].FileID)
		if err != nil {
			return err
		}

		err = b.updateParticipantInDB("photo", imageUrl, message.From.ID)
		if err != nil {
			return err
		} else {
			msg := tgbotapi.NewMessage(message.Chat.ID, "Your photo successfully updated!")
			msg.ReplyMarkup = registrationKeyboard
			_, err := b.bot.Send(msg)
			return err
		}
	case "description":
		if err := b.deleteCache(message.From.ID); err != nil {
			return err
		}

		err := b.updateParticipantInDB("information", message.Text, message.From.ID)
		if err != nil {
			return err
		} else {
			msg := tgbotapi.NewMessage(message.Chat.ID, "Your description successfully updated!")
			msg.ReplyMarkup = registrationKeyboard
			_, err := b.bot.Send(msg)
			return err
		}
	default:
		msg := tgbotapi.NewMessage(message.Chat.ID, "error occured while getting cache, use /start to resolve problem...")
		_, err := b.bot.Send(msg)
		return err
	}
}
