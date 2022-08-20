package bot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

const (
	// command list
	start   = "start"
	support = "support"

	// choose role
	register = "Register!"
	vote     = "Vote!"

	// register
	name        = "Change a Name"
	photo       = "Add a Photo"
	description = "Write a Description"
	profile     = "Show my Profile!"
	delete      = "Delete my Profile!"
	back        = "🔙"

	// votes
	like    = "👍"
	dislike = "👎"
)

var roleKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(register),
		tgbotapi.NewKeyboardButton(vote),
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
