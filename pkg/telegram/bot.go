package telegram

import (
	"coloriAI/internal/config"
	"coloriAI/pkg/logger"
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
)

type Bot struct {
	Api           tgbotapi.BotAPI
	Logger        *zap.Logger
	isDebug       bool
	updateTimeout int
	ctx           context.Context
}

func NewBot(cfg config.BotConfig, logger *zap.Logger) (*Bot, error) {

	botApi, err := tgbotapi.NewBotAPI(cfg.BotToken)
	if err != nil {
		return nil, err
	}

	botApi.Debug = cfg.DebugMode

	return &Bot{
		Api:           *botApi,
		isDebug:       cfg.DebugMode,
		updateTimeout: cfg.UpdateBotTimeout,
		Logger:        logger.Named(cfg.BotName),
	}, nil
}

func (b *Bot) SetLogger() error {
	loggerAdapter := logger.NewAdapter(b.Logger)

	err := tgbotapi.SetLogger(loggerAdapter)
	if err != nil {
		b.Logger.Error("Failed to set logger", zap.Error(err))
		return err
	}

	return nil
}

func (b *Bot) GetUpdateConfig() tgbotapi.UpdateConfig {
	updateConfig := tgbotapi.NewUpdate(0)

	updateConfig.Timeout = b.updateTimeout

	return updateConfig
}
