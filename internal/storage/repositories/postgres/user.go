package storage

import (
	"coloriAI/internal/entities/user"
	"coloriAI/pkg/postgres"
	"context"
)

// implementation for postgres

type UserRepository struct {
	db *postgres.Database
}

func (r *UserRepository) CreateUser(ctx context.Context, user user.User) (error, string) {
	row := r.db.Pool.QueryRow(ctx, "INSERT INTO users (id, name, telegram_id) VALUES ($1, $2, $3)", user.ID, user.Name, user.TelegramID)

	var userId string

	err := row.Scan(&user.ID)
	if err != nil {
		return err, ""
	}

	return nil, userId
}

func (r *UserRepository) GetUser(ctx context.Context, id int) (user.User, error) {
	//TODO implement me
	panic("implement me")
}

func (r *UserRepository) GetUserByID(ctx context.Context, id string) (user.User, error) {
	//TODO implement me
	panic("implement me")
}

func (r *UserRepository) GetUserByTelegramID(ctx context.Context, id int) (user.User, error) {
	//TODO implement me
	panic("implement me")
}

func (r *UserRepository) GetUsers(ctx context.Context) ([]user.User, error) {
	//TODO implement me
	panic("implement me")
}

func NewUserRepository(db *postgres.Database) *UserRepository {
	return &UserRepository{db: db}
}
