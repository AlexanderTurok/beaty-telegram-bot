package bot

import (
	"fmt"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (b *Bot) handleMessages(message *tgbotapi.Message) error {
	switch message.Text {
	case register:
		exists, err := b.service.ParticipantData.IsParticipant(message.From.ID)
		if err != nil {
			return err
		}
		if !exists {
			err := b.service.ParticipantData.AddParticipant(message.From.ID)
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

		err = b.service.ParticipantCache.SetCache(message.From.ID, "name")
		return err
	case photo:
		msg := tgbotapi.NewMessage(message.Chat.ID, "Waiting for you photo...")
		_, err := b.bot.Send(msg)
		if err != nil {
			return err
		}

		err = b.service.ParticipantCache.SetCache(message.From.ID, "photo")
		return err
	case description:
		msg := tgbotapi.NewMessage(message.Chat.ID, "Describe yourself")
		_, err := b.bot.Send(msg)
		if err != nil {
			return err
		}

		err = b.service.ParticipantCache.SetCache(message.From.ID, "description")
		return err
	case profile:
		user, err := b.service.ParticipantData.GetParticipant(message.From.ID)
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
		err := b.service.ParticipantData.DeleteParticipant(message.From.ID)
		if err != nil {
			return err
		}

		msg := tgbotapi.NewMessage(message.Chat.ID, "Your profile successfully deleted! Choose avaible method: ")
		msg.ReplyMarkup = roleKeyboard
		_, err = b.bot.Send(msg)
		if err != nil {
			return err
		}

		err = b.service.ParticipantCache.DeleteCache(message.From.ID)
		return err
	case back:
		if err := b.service.ParticipantCache.DeleteCache(message.From.ID); err != nil {
			return err
		}

		msg := tgbotapi.NewMessage(message.Chat.ID, "Choose avaible method: ")
		msg.ReplyMarkup = roleKeyboard
		_, err := b.bot.Send(msg)

		return err
	case vote:
		participantID, err := b.service.VoterCache.GetCache(message.From.ID)
		if err != nil {
			return err
		}

		if participantID == "" {
			err := b.service.VoterCache.SetCache(message.From.ID, "0")
			if err != nil {
				return err
			}
		}

		p, err := b.service.ParticipantData.GetAllParticipants()
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
			if err := b.service.VoterCache.SetCache(message.From.ID, fmt.Sprint(id+1)); err != nil {
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
		participantID, err := b.service.VoterCache.GetCache(message.From.ID)
		if err != nil {
			return err
		}
		b.service.ParticipantData.UpdateParticipant("uuid", participantID, message.From.ID)

		p, err := b.service.ParticipantData.GetAllParticipants()
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
			if err := b.service.VoterCache.SetCache(message.From.ID, fmt.Sprint(id)); err != nil {
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
		participantID, err := b.service.VoterCache.GetCache(message.From.ID)
		if err != nil {
			return err
		}

		p, err := b.service.ParticipantData.GetAllParticipants()
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
			err = b.service.VoterCache.SetCache(message.From.ID, fmt.Sprint(id))
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
