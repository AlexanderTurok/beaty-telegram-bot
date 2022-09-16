package bot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func (b *Bot) handleError(userId int64, err error) error {
	msg := tgbotapi.NewMessage(userId, err.Error())
	_, err = b.bot.Send(msg)

	return err
}
