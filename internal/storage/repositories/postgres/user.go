package storage

import (
	"coloriAI/internal/entities/user"
	"coloriAI/pkg/postgres"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
)

// implementation for postgres

type UserRepository struct {
	db *postgres.Database
}

func (r *UserRepository) CreateUser(ctx context.Context, user user.User) (error, string) {
	row := r.db.Pool.QueryRow(ctx, "INSERT INTO users (username, telegram_id) VALUES ($1, $2)", user.Username, user.TelegramID)

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

func (r *UserRepository) GetUserByTelegramID(ctx context.Context, id int64) (user.User, error) {
	row := r.db.Pool.QueryRow(ctx, "SELECT * FROM users WHERE telegram_id = $1", id)

	var u user.User
	err := row.Scan(&u.ID, &u.TelegramID, &u.Username, &u.CreatedAt, &u.UpdatedAt)

	if err != nil {
		if err == pgx.ErrNoRows {
			fmt.Printf("No u found with telegram_id %d\n", id)
			return u, nil // Возвращаем пустого пользователя, потому что данных нет
		} else {
			return u, fmt.Errorf("failed to scan u: %v", err)
		}
	}

	return u, nil
}

func (r *UserRepository) GetUsers(ctx context.Context) ([]user.User, error) {
	//TODO implement me
	panic("implement me")
}

func NewUserRepository(db *postgres.Database) *UserRepository {
	return &UserRepository{db: db}
}
