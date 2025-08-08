package services

import (
	"coloriAI/internal/entities/user"
	services "coloriAI/internal/services/users"
	"coloriAI/internal/storage"
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
)

type UserService interface {
	CreateUser(ctx context.Context, user *tgbotapi.User) (error, string)
	FindUserByTelegramID(ctx context.Context, id int64) (user.User, error)
}

type Service struct {
	UserService UserService
}

func NewService(repository *storage.Repository, logger *zap.Logger) *Service {
	return &Service{
		UserService: services.NewUserService(repository, logger),
	}
}
