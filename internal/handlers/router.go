package handlers

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Router struct {
	commandHandler *CommandHandler
}

func NewRouter(cmd *CommandHandler) *Router {
	return &Router{
		commandHandler: cmd,
	}
}

func (r *Router) HandleUpdate(ctx context.Context, update tgbotapi.Update) {
	if update.Message != nil && update.Message.IsCommand() {
		r.commandHandler.Handle(ctx, update.Message)
		return
	}

	if update.CallbackQuery != nil {
		panic("implement me")
	}
}
