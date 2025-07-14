ifneq (,$(wildcard ./.env))
    include .env
    export
endif


lint:
	golangci-lint run ./...

migrations-up:
	goose -dir $(MIGRATION_DIR) postgres $(DB_URL) up

migrations-status:
	goose -dir $(MIGRATION_DIR) postgres $(DB_URL) status

migrations-down:
	goose -dir $(MIGRATION_DIR) postgres $(DB_URL) down
