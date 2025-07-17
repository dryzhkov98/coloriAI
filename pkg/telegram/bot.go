package telegram

import (
	"coloriAI/internal/config"
	"coloriAI/pkg/logger"
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
)

type Bot struct {
	api           tgbotapi.BotAPI
	isDebug       bool
	updateTimeout int
	ctx           context.Context
}

func NewBot(cfg config.BotConfig) (*Bot, error) {

	bot, err := tgbotapi.NewBotAPI(cfg.BotToken)
	if err != nil {
		return nil, err
	}

	bot.Debug = cfg.DebugMode

	return &Bot{
		api:           *bot,
		isDebug:       cfg.DebugMode,
		updateTimeout: cfg.UpdateBotTimeout,
	}, nil
}

func (b *Bot) Start() error {
	appLogger := logger.Get()

	err := tgbotapi.SetLogger(&logger.ZapLoggerAdapter{})
	if err != nil {
		appLogger.Error("Failed to set logger", zap.Error(err))
		return err
	}
	updateConfig := tgbotapi.NewUpdate(0)

	updateConfig.Timeout = b.updateTimeout

	updates := b.api.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message != nil {
			_, err := b.api.Send(tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text))
			if err != nil {
				panic(err)
			}
		}
	}

	return nil
}
