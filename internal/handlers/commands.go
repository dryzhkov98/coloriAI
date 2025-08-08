package handlers

import (
	"coloriAI/pkg/telegram"
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
)

func (h *Handlers) Command(ctx context.Context, bot *telegram.Bot, update tgbotapi.Update) {
	switch update.Message.Command() {
	case "start":
		err := h.StartHandler.HandleCommand(ctx, bot, update)
		if err != nil {
			bot.Logger.Error("Error while handling command", zap.Error(err))
		}
	}
}
