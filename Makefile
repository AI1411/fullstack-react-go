# PostgreSQL connection settings
DB_HOST ?= db
DB_PORT ?= 5432
DB_USER ?= postgres
DB_PASSWORD ?= postgres
DB_NAME ?= gen
DB_SSLMODE ?= disable

# Construct database URL
DATABASE_URL = postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)

.PHONY: migrate migrate-up migrate-down migrate-version migrate-create

# Run migration up
migrate:
	docker compose exec migration migrate -source file://./ -database '$(DATABASE_URL)' up

# Alias for migrate
migrate-up: migrate

# Run migration down
migrate-down:
	docker compose exec migration migrate -source file://./ -database '$(DATABASE_URL)' down 1

# Show current migration version
migrate-version:
	docker compose exec migration migrate -source file://./ -database '$(DATABASE_URL)' version

# Create new migration file
migrate-create:
	@read -p "Enter migration name: " name; \
	docker compose exec migration migrate create -ext sql -dir ./ -seq $$name

logs:
	docker logs gen-api -f --tail 100