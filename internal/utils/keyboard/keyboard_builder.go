package keyboardbuilder

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Keyboard struct {
	rows [][]tgbotapi.KeyboardButton
}

type KeyboardBuilder interface {
	AddRow(texts ...string) KeyboardBuilder
	Build() tgbotapi.ReplyKeyboardMarkup
}

func (b *Keyboard) AddRow(texts ...string) KeyboardBuilder {
	var row []tgbotapi.KeyboardButton

	for _, text := range texts {
		row = append(row, tgbotapi.NewKeyboardButton(text))
	}

	b.rows = append(b.rows, row)
	return b
}
func (b *Keyboard) Build() tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.ReplyKeyboardMarkup{
		Keyboard:       b.rows,
		ResizeKeyboard: true,
	}
}
