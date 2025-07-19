package keyboard

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type ReplyBuilder struct {
	rows [][]tgbotapi.KeyboardButton
}

func NewReply() *ReplyBuilder {
	return &ReplyBuilder{}
}

func (b *ReplyBuilder) Row(labels ...string) *ReplyBuilder {
	var row []tgbotapi.KeyboardButton

	for _, label := range labels {
		row = append(row, tgbotapi.NewKeyboardButton(label))
	}
	b.rows = append(b.rows, row)
	return b
}

func (b *ReplyBuilder) Build() tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.ReplyKeyboardMarkup{
		Keyboard:       b.rows,
		ResizeKeyboard: true,
	}
}
