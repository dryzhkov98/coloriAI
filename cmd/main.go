package main

import (
	"coloriAI/internal/config"
	"coloriAI/pkg/logger"
	"coloriAI/pkg/postgres"
	"coloriAI/pkg/telegram"
	"context"
	"time"
)

func main() {
	ctx := context.Background()

	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	logger.New(cfg)

	appLogger := logger.Get()

	db, err := postgres.NewDatabase(ctx, cfg.DBConfig, appLogger)
	if err != nil {
		panic(err.Error())
	}
	defer db.Pool.Close()

	bot, err := telegram.NewBot(cfg.BotConfig)
	if err != nil {
		panic(err.Error())
	}

	go func() {
		err := bot.Start()
		if err != nil {
			panic(err)
		}
	}()
	appLogger.Info("Bot is started")
	<-ctx.Done()
	time.Sleep(2 * time.Second)

}
