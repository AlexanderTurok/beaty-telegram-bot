package telegram

import (
	"fmt"
	"strconv"

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
	description = "Write a Description"
	profile     = "Show my Profile!"
	delete      = "Delete my Profile!"
	back        = "Go Back🔙"

	// votes
	like    = "👍"
	dislike = "👎"
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
		tgbotapi.NewKeyboardButton(photo),
		tgbotapi.NewKeyboardButton(description),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(profile),
		tgbotapi.NewKeyboardButton(delete),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(back),
	),
)

var voteKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(like),
		tgbotapi.NewKeyboardButton(dislike),
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
		}
		msg := tgbotapi.NewMessage(message.Chat.ID, "Here is available methods:")
		msg.ReplyMarkup = registrationKeyboard
		_, err = b.bot.Send(msg)
		return err
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
	case profile:
		user, err := b.getParticipantFromDB("uuid", message.From.ID)
		fmt.Println(user.Photo, user.Nickname, user.Information)
		if err != nil {
			return err
		}

		msg := tgbotapi.NewPhotoShare(message.Chat.ID, user.Photo)
		msg.Caption = fmt.Sprintf("%s, %s", user.Nickname, user.Information)
		_, err = b.bot.Send(msg)

		return err
	case delete:
		err := b.deleteParticipantFromDB(message.From.ID)
		if err != nil {
			return err
		}

		msg := tgbotapi.NewMessage(message.Chat.ID, "Your profile successfully deleted! Choose avaible method: ")
		msg.ReplyMarkup = roleKeyboard
		_, err = b.bot.Send(msg)
		if err != nil {
			return err
		}

		err = b.deleteCache(message.From.ID)
		return err
	case back:
		if err := b.deleteCache(message.From.ID); err != nil {
			return err
		}

		msg := tgbotapi.NewMessage(message.Chat.ID, "Choose avaible method: ")
		msg.ReplyMarkup = roleKeyboard
		_, err := b.bot.Send(msg)

		return err
	case voter:
		participantID := b.getCache(message.From.ID)
		if participantID == "" {
			err := b.setCache(message.From.ID, 0)
			if err != nil {
				return err
			}
		}

		p, err := b.getAllParticipantsFromDB()
		if err != nil {
			return err
		}

		id, _ := strconv.Atoi(participantID) // convert string to int

		if id >= len(*p) {
			msg := tgbotapi.NewMessage(message.Chat.ID, "You voted for all participants. Wait some time for new participants...")
			msg.ReplyMarkup = roleKeyboard
			_, err := b.bot.Send(msg)
			if err != nil {
				return err
			}
		} else {
			if err := b.setCache(message.From.ID, id+1); err != nil {
				return err
			}

			msg := tgbotapi.NewPhotoShare(message.Chat.ID, (*p)[id].Photo)
			msg.ReplyMarkup = tgbotapi.ReplyKeyboardMarkup{
				Keyboard:        voteKeyboard.Keyboard,
				OneTimeKeyboard: true,
				ResizeKeyboard:  true,
			}
			msg.Caption = fmt.Sprintf("%s, %s", (*p)[id].Nickname, (*p)[id].Information)
			_, err = b.bot.Send(msg)
			if err != nil {
				return err
			}
		}
		return nil
	case like:
		participantID := b.getCache(message.From.ID)
		b.updateVotesInDB(participantID)

		p, err := b.getAllParticipantsFromDB()
		if err != nil {
			return err
		}

		id, _ := strconv.Atoi(participantID)
		id += 1

		if id >= len(*p) {
			msg := tgbotapi.NewMessage(message.Chat.ID, "You voted for all participants. Wait some time for new participants...")
			msg.ReplyMarkup = roleKeyboard
			_, err := b.bot.Send(msg)
			if err != nil {
				return err
			}
		} else {
			if err := b.setCache(message.From.ID, id); err != nil {
				return err
			}

			msg := tgbotapi.NewPhotoShare(message.Chat.ID, (*p)[id].Photo)
			msg.ReplyMarkup = tgbotapi.ReplyKeyboardMarkup{
				Keyboard:        voteKeyboard.Keyboard,
				OneTimeKeyboard: true,
				ResizeKeyboard:  true,
			}
			msg.Caption = fmt.Sprintf("%s, %s", (*p)[id].Nickname, (*p)[id].Information)
			_, err = b.bot.Send(msg)

			return err
		}

		return nil
	case dislike:
		participantID := b.getCache(message.From.ID)

		p, err := b.getAllParticipantsFromDB()
		if err != nil {
			return err
		}

		id, _ := strconv.Atoi(participantID)
		id += 1

		if id >= len(*p) {
			msg := tgbotapi.NewMessage(message.Chat.ID, "You voted for all participants. Wait some time for new participants...")
			msg.ReplyMarkup = roleKeyboard
			_, err := b.bot.Send(msg)
			if err != nil {
				return err
			}
		} else {
			err = b.setCache(message.From.ID, id)
			if err != nil {
				return err
			}
			msg := tgbotapi.NewPhotoShare(message.Chat.ID, (*p)[id].Photo)
			msg.ReplyMarkup = tgbotapi.ReplyKeyboardMarkup{
				Keyboard:        voteKeyboard.Keyboard,
				OneTimeKeyboard: true,
				ResizeKeyboard:  true,
			}
			msg.Caption = fmt.Sprintf("%s, %s", (*p)[id].Nickname, (*p)[id].Information)
			_, err = b.bot.Send(msg)

			return err
		}
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

		err := b.updateParticipantInDB("photo", (*message.Photo)[0].FileID, message.From.ID)
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
		msg := tgbotapi.NewMessage(message.Chat.ID, "something went from with cache handler")
		_, err := b.bot.Send(msg)
		return err
	}
}
