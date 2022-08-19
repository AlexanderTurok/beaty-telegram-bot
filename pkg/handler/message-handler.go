package bot

import (
	"fmt"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (b *Bot) handleMessages(message *tgbotapi.Message) error {
	switch message.Text {
	case register:
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
		if user.Photo != "" && user.Nickname != "" && user.Information != "" {
			_, err = b.bot.Send(msg)
		}

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
	case vote:
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
			if (*p)[id].Photo != "" && (*p)[id].Nickname != "" && (*p)[id].Information != "" {
				_, err = b.bot.Send(msg)
			}

			return err
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
			if (*p)[id].Photo != "" && (*p)[id].Nickname != "" && (*p)[id].Information != "" {
				_, err = b.bot.Send(msg)
			}
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
			if (*p)[id].Photo != "" && (*p)[id].Nickname != "" && (*p)[id].Information != "" {
				_, err = b.bot.Send(msg)
			}

			return err
		}
		return nil
	default:
		msg := tgbotapi.NewMessage(message.Chat.ID, "unknown message...")
		_, err := b.bot.Send(msg)
		return err
	}
}
