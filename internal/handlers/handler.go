package handlers

import (
	handlers "coloriAI/internal/handlers/start"
	"coloriAI/internal/services"
	"coloriAI/pkg/telegram"
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type StartHandler interface {
	HandleCommand(ctx context.Context, bot telegram.Bot, update tgbotapi.Update) error
}

type Handlers struct {
	StartHandler *handlers.StartHandler
}

func NewHandler(userService services.UserService) *Handlers {
	return &Handlers{
		StartHandler: handlers.NewStartHandler(userService),
	}
}
