-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
    id uuid PRIMARY KEY UNIQUE NOT NULL DEFAULT gen_random_uuid(),
    telegram_id bigint UNIQUE NOT NULL,
    username TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
    DROP TABLE IF EXISTS users;
-- +goose StatementEnd
