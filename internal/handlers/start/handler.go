package handlers

import (
	"coloriAI/internal/services"
	"coloriAI/pkg/telegram"
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type StartHandler struct {
	userService services.UserService
}

func NewStartHandler(userService services.UserService) *StartHandler {
	return &StartHandler{
		userService: userService,
	}
}

func (h *StartHandler) HandleCommand(ctx context.Context, bot telegram.Bot, update tgbotapi.Update) error {
	panic("implement me")
}
