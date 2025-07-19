package storage

import (
	"coloriAI/internal/entities/user"
	storage "coloriAI/internal/storage/repositories/postgres"
	"coloriAI/pkg/postgres"
	"context"
)

// Repository layer abstraction

type UserRepository interface {
	CreateUser(ctx context.Context, user user.User) (error, string)
	GetUser(ctx context.Context, id int) (user.User, error)
	GetUserByID(ctx context.Context, id string) (user.User, error)
	GetUserByTelegramID(ctx context.Context, id int) (user.User, error)
	GetUsers(ctx context.Context) ([]user.User, error)
}

type FoodJournalRepository interface {
	CreateFoodJournal(ctx context.Context, userID string) error
}

type Repository struct {
	UserRepository        UserRepository
	FoodJournalRepository FoodJournalRepository
}

func NewRepository(db *postgres.Database) *Repository {
	return &Repository{
		UserRepository: storage.NewUserRepository(db),
	}
}
