package services

import (
	"coloriAI/internal/entities/user"
	"context"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (us *UserService) CreateUser(ctx context.Context, user user.User) (error, string) {
	panic("implement me")
}
