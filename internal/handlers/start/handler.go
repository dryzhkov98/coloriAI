package handlers

import (
	"coloriAI/internal/entities/user"
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

func (h *StartHandler) HandleCommand(ctx context.Context, bot *telegram.Bot, update tgbotapi.Update) error {
	logger := bot.Logger
	logger.Info("Start command")

	existing, err := h.userService.FindUserByTelegramID(ctx, update.Message.From.ID)
	if err != nil {
		return err
	}

	if existing != (user.User{}) {
		logger.Info("User already exists")
		return nil
	}

	err, _ = h.userService.CreateUser(ctx, update.Message.From)

	if err != nil {
		return err
	}
	return nil

}
