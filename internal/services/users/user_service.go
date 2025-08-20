package services

import (
	"coloriAI/internal/entities/user"
	"coloriAI/internal/storage"
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
)

type UserService struct {
	repository *storage.Repository
	logger     *zap.Logger
}

func (us *UserService) FindUserByTelegramID(ctx context.Context, id int64) (user.User, error) {
	existedUser, err := us.repository.UserRepository.GetUserByTelegramID(ctx, id)
	if err != nil {
		us.logger.Error("Error while getting user by telegram id", zap.Error(err))
		return user.User{}, err
	}

	return existedUser, nil
}

func NewUserService(repository *storage.Repository, logger *zap.Logger) *UserService {
	return &UserService{
		repository: repository,
		logger:     logger,
	}
}

func (us *UserService) CreateUser(ctx context.Context, tgUser *tgbotapi.User) (error, string) {
	user := user.User{
		TelegramID: tgUser.ID,
		Username:   tgUser.FirstName,
	}
	err, id := us.repository.UserRepository.CreateUser(ctx, user)
	if err != nil {
		us.logger.Error("Error while creating user", zap.Error(err))
		return err, ""
	}

	us.logger.Info("User created successfully", zap.String("id", id))
	return nil, id
}
