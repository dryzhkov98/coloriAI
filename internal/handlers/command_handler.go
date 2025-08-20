package handlers

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type CommandHandler struct {
	startUC *interface{}
	helpUC  *interface{}
}

type CmdHandler interface {
	Handle(ctx context.Context, cmd *tgbotapi.Message)
}

func NewCommandHandler() *CommandHandler {
	return &CommandHandler{}
}

func (c *CommandHandler) Handle(ctx context.Context, cmd *tgbotapi.Message) {
	switch cmd.Command() {
	case "start":
		panic("implement me")

	case "help":
		panic("implement me")

	default:
		panic("unknown command")
	}
}
