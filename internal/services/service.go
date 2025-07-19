package services

import (
	"coloriAI/internal/entities/user"
	services "coloriAI/internal/services/users"
	"context"
)

type UserService interface {
	CreateUser(ctx context.Context, user user.User) (error, string)
}

type Service struct {
	UserService UserService
}

func NewService() *Service {
	return &Service{
		UserService: services.NewUserService(),
	}
}
