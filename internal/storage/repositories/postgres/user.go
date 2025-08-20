package storage

import (
	"coloriAI/internal/entities/user"
	"coloriAI/pkg/postgres"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
)

// implementation for postgres

type UserRepository struct {
	db *postgres.Database
}

func (r *UserRepository) CreateUser(ctx context.Context, user user.User) (error, string) {
	row := r.db.Pool.QueryRow(ctx, "INSERT INTO users (username, telegram_id) VALUES ($1, $2) RETURNING id", user.Username, user.TelegramID)

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

func (r *UserRepository) GetUserByTelegramID(ctx context.Context, tgID int64) (user.User, error) {
	var u user.User

	err := r.db.Pool.QueryRow(ctx,
		`SELECT id, username, telegram_id, created_at, updated_at
		   FROM users
		  WHERE telegram_id = $1`,
		tgID,
	).Scan(&u.ID, &u.Username, &u.TelegramID, &u.CreatedAt, &u.UpdatedAt)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return user.User{}, pgx.ErrNoRows
		}
		return user.User{}, fmt.Errorf("get user by telegram_id=%d: %w", tgID, err)
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
