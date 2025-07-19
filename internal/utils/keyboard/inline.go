package keyboard

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type InlineBuilder struct {
	rows [][]tgbotapi.InlineKeyboardButton
}

type InlineButton struct {
	Label        string
	CallbackData string
}

func NewInline() *InlineBuilder {
	return &InlineBuilder{}
}

func (b *InlineBuilder) Row(buttons ...InlineButton) *InlineBuilder {
	var row []tgbotapi.InlineKeyboardButton

	for _, btn := range buttons {
		row = append(row, tgbotapi.NewInlineKeyboardButtonData(btn.Label, btn.CallbackData))
	}
	b.rows = append(b.rows, row)
	return b
}

func (b *InlineBuilder) Build() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(b.rows...)
}
