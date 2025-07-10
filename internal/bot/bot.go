package bot

import (
	"coloriAI/internal/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	api           tgbotapi.BotAPI
	isDebug       bool
	updateTimeout int
}

func New(cfg *config.Config) (*Bot, error) {
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

func (b *Bot) Start() {
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = b.updateTimeout

	updates := b.api.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message != nil {
			err := b.handleCommand(update.Message)
			if err != nil {
				return
			}
		}
	}
}
