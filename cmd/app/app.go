package app

import (
	"coloriAI/internal/config"
	"coloriAI/internal/handlers"
	"coloriAI/internal/services"
	"coloriAI/internal/storage"
	"coloriAI/pkg/logger"
	"coloriAI/pkg/postgres"
	"coloriAI/pkg/telegram"
	"context"
	"go.uber.org/zap"
	"time"
)

type App struct {
	bot          *telegram.Bot
	appConfig    *config.Config
	db           *postgres.Database
	repositories *storage.Repository
	services     *services.Service
	handlers     *handlers.Handlers
	ctx          context.Context
	logger       *zap.Logger
}

func NewApp(cfg *config.Config) *App {
	ctx := context.Background()

	appLogger := logger.New(cfg)

	app := &App{
		appConfig: cfg,
		ctx:       ctx,
		logger:    appLogger,
	}

	app.mustCreateDB()
	app.mustCreateBot()

	app.createRepositories()
	app.createServices()

	return app

}

func (a *App) Run() {
	go a.startBot()
	a.logger.Info("Bot is started")
	<-a.ctx.Done()
	a.logger.Info("Bot is stopped")
	time.Sleep(2 * time.Second)
}

func (a *App) mustCreateDB() {
	db, err := postgres.NewDatabase(a.ctx, a.appConfig.DBConfig, a.logger)
	if err != nil {
		a.logger.Error("Error while creating DB", zap.Error(err))
		panic(err.Error())
	}

	a.db = db
}

func (a *App) mustCreateBot() {
	bot, err := telegram.NewBot(a.appConfig.BotConfig, a.logger)
	if err != nil {
		a.logger.Error("Error while creating bot", zap.Error(err))
		panic(err.Error())
	}

	a.bot = bot

}

func (a *App) startBot() {
	a.logger.Info("Starting bot", zap.String("bot_name", a.appConfig.BotConfig.BotName))

	updateConfig := a.bot.GetUpdateConfig()

	updates := a.bot.Api.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message.IsCommand() {
		}
	}

}

func (a *App) createRepositories() {
	a.logger.Info("Creating repositories")
	repositories := storage.NewRepository(a.db)

	a.repositories = repositories
}

func (a *App) createServices() {
	a.logger.Info("Creating services")

	service := services.NewService()

	a.services = service
}
func (a *App) createHandlers() {
	a.logger.Info("Creating handlers")

	handler := handlers.NewHandler(a.services.UserService)

	a.handlers = handler
}
