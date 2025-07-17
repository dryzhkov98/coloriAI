package storage

import (
	"coloriAI/internal/entities/user"
	"coloriAI/pkg/postgres"
	"context"
)

// implementation for postgres

type UserRepository struct {
	db postgres.Database
}

func NewUserRepository(db postgres.Database) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(ctx context.Context, args *user.User) (error, string) {
	row := r.db.Pool.QueryRow(ctx, "INSERT INTO users (id, name, telegram_id) VALUES ($1, $2, $3)", args.ID, args.Name, args.TelegramID)

	var userId string

	err := row.Scan(&args.ID)
	if err != nil {
		return err, ""
	}

	return nil, userId

}
