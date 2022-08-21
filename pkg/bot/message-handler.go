package bot

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (b *Bot) handleMessages(message *tgbotapi.Message) error {
	switch message.Text {
	case register:
		return b.handleRegistration(message)
	case name:
		return b.getName(message)
	case photo:
		return b.getPhoto(message)
	case description:
		return b.getDescription(message)
	case profile:
		return b.getProfile(message.From.ID, message, false)
	case delete:
		return b.deleteProfile(message)
	case back:
		return b.back(message)
	case vote:
		return b.handleVote(message)
	case like:
		return b.handleLike(message)
	case dislike:
		return b.handleDislike(message)
	default:
		return b.handleDefaultMessage(message)
	}
}

func (b *Bot) handleRegistration(message *tgbotapi.Message) error {
	if err := b.service.Participant.Register(message); err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, "Here is available methods:")
	msg.ReplyMarkup = registrationKeyboard
	_, err := b.bot.Send(msg)

	return err
}

func (b *Bot) getName(message *tgbotapi.Message) error {
	if err := b.service.Participant.SetCache(message.From.ID, "name"); err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, "Send your name visible to others")
	_, err := b.bot.Send(msg)

	return err
}

func (b *Bot) getPhoto(message *tgbotapi.Message) error {
	if err := b.service.Participant.SetCache(message.From.ID, "photo"); err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, "Waiting for you photo...")
	_, err := b.bot.Send(msg)

	return err
}

func (b *Bot) getDescription(message *tgbotapi.Message) error {
	if err := b.service.Participant.SetCache(message.From.ID, "description"); err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, "Describe yourself")
	_, err := b.bot.Send(msg)

	return err
}

func (b *Bot) deleteProfile(message *tgbotapi.Message) error {
	if err := b.service.Participant.DeleteParticipant(message.From.ID); err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, "Your profile successfully deleted! Choose avaible method: ")
	msg.ReplyMarkup = roleKeyboard
	_, err := b.bot.Send(msg)

	return err
}

func (b *Bot) back(message *tgbotapi.Message) error {
	if err := b.service.Participant.DeleteCache(message.From.ID); err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, "Choose avaible method: ")
	msg.ReplyMarkup = roleKeyboard
	_, err := b.bot.Send(msg)

	return err
}

func (b *Bot) handleVote(message *tgbotapi.Message) error {
	id, err := b.service.Voter.GetID(message)
	if err != nil {
		return err
	}

	p, err := b.service.Participant.GetAllParticipants()
	if err != nil {
		return err
	}

	if id >= len(*p) {
		return b.handleEndOfParticipants(message)
	} else {
		if err := b.service.Voter.SetCache(message.From.ID, fmt.Sprint(id+1)); err != nil {
			return err
		}

		if err := b.getProfile((*p)[id].Id, message, true); err != nil {
			return err
		}
	}

	return nil
}

func (b *Bot) handleLike(message *tgbotapi.Message) error {
	id, err := b.service.GetID(message)
	if err != nil {
		return err
	}

	if err := b.service.Participant.UpdateParticipant("uuid", fmt.Sprint(id), message.From.ID); err != nil {
		return err
	}

	p, err := b.service.Participant.GetAllParticipants()
	if err != nil {
		return err
	}

	id += 1

	if id >= len(*p) {
		return b.handleEndOfParticipants(message)
	} else {
		if err := b.service.Voter.SetCache(message.From.ID, fmt.Sprint(id)); err != nil {
			return err
		}

		if err := b.getProfile((*p)[id].Id, message, true); err != nil {
			return err
		}
	}

	return nil
}

func (b *Bot) handleDislike(message *tgbotapi.Message) error {
	id, err := b.service.GetID(message)
	if err != nil {
		return err
	}

	p, err := b.service.Participant.GetAllParticipants()
	if err != nil {
		return err
	}

	id += 1

	if id >= len(*p) {
		return b.handleEndOfParticipants(message)
	} else {
		if err := b.service.Voter.SetCache(message.From.ID, fmt.Sprint(id)); err != nil {
			return err
		}

		if err := b.getProfile((*p)[id].Id, message, true); err != nil {
			return err
		}
	}

	return nil
}

func (b *Bot) handleDefaultMessage(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "unknown message...")
	_, err := b.bot.Send(msg)

	return err
}

func (b *Bot) handleEndOfParticipants(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "You voted for all participants. Wait some time for new participants...")
	msg.ReplyMarkup = roleKeyboard
	_, err := b.bot.Send(msg)

	return err
}

func (b *Bot) getProfile(participantID int, message *tgbotapi.Message, showKeyboard bool) error {
	user, err := b.service.Participant.GetParticipant(participantID)
	if err != nil {
		return err
	}

	msg := tgbotapi.NewPhotoShare(message.Chat.ID, user.Photo)
	msg.Caption = fmt.Sprintf("%s, %s", user.Nickname, user.Information)
	if showKeyboard {
		msg.ReplyMarkup = tgbotapi.ReplyKeyboardMarkup{
			Keyboard:        voteKeyboard.Keyboard,
			OneTimeKeyboard: true,
			ResizeKeyboard:  true,
		}
	}
	if user.Photo != "" && user.Nickname != "" && user.Information != "" {
		_, err = b.bot.Send(msg)
	}

	return err
}
