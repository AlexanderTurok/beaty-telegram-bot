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
		return b.getProfile(message)
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
	if err := b.service.Participant.Register(message.Chat.ID); err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, "Here is available methods:")
	msg.ReplyMarkup = registrationKeyboard
	_, err := b.bot.Send(msg)

	return err
}

func (b *Bot) getName(message *tgbotapi.Message) error {
	if err := b.service.Participant.GetName(message.Chat.ID, nameCache); err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, "Send your name visible to others")
	_, err := b.bot.Send(msg)

	return err
}

func (b *Bot) getPhoto(message *tgbotapi.Message) error {
	if err := b.service.Participant.GetPhoto(message.Chat.ID, photoCache); err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, "Waiting for you photo...")
	_, err := b.bot.Send(msg)

	return err
}

func (b *Bot) getDescription(message *tgbotapi.Message) error {
	if err := b.service.Participant.GetDescription(message.Chat.ID, descriptionCache); err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, "Describe yourself")
	_, err := b.bot.Send(msg)

	return err
}

func (b *Bot) getProfile(message *tgbotapi.Message) error {
	user, err := b.service.Participant.Get(message.Chat.ID)
	if err != nil {
		b.handleError(message.Chat.ID, err)
		return err
	}

	msg := tgbotapi.NewPhotoShare(message.Chat.ID, fmt.Sprint(user.Photo))
	msg.Caption = fmt.Sprintf("%s, %s", user.Name, user.Description)
	msg.ReplyMarkup = registrationKeyboard
	_, err = b.bot.Send(msg)

	return err
}

func (b *Bot) deleteProfile(message *tgbotapi.Message) error {
	if err := b.service.Participant.Delete(message.Chat.ID); err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, "Your profile successfully deleted! Choose avaible method: ")
	msg.ReplyMarkup = roleKeyboard
	_, err := b.bot.Send(msg)

	return err
}

func (b *Bot) back(message *tgbotapi.Message) error {
	if err := b.service.Participant.DeleteCache(message.Chat.ID); err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, "Choose avaible method: ")
	msg.ReplyMarkup = roleKeyboard
	_, err := b.bot.Send(msg)

	return err
}

func (b *Bot) handleVote(message *tgbotapi.Message) error {
	if err := b.service.Voter.Create(message.Chat.ID); err != nil {
		return err
	}

	participant, err := b.service.Voter.GetParticipant(message.Chat.ID)
	if err != nil {
		b.handleError(message.Chat.ID, err)
		return err
	}

	msg := tgbotapi.NewPhotoShare(message.Chat.ID, fmt.Sprint(participant.Photo))
	msg.Caption = fmt.Sprintf("%s, %s", participant.Name, participant.Description)
	msg.ReplyMarkup = voteKeyboard
	_, err = b.bot.Send(msg)

	return err
}

func (b *Bot) handleLike(message *tgbotapi.Message) error {
	return nil
}

func (b *Bot) handleDislike(message *tgbotapi.Message) error {
	return nil
}

func (b *Bot) handleDefaultMessage(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "unknown message...")
	_, err := b.bot.Send(msg)

	return err
}
