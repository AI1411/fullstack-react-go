# PostgreSQL connection settings
DB_HOST ?= db
DB_PORT ?= 5432
DB_USER ?= postgres
DB_PASSWORD ?= postgres
DB_NAME ?= gen
DB_SSLMODE ?= disable
# PostgreSQL connection settings for test database
DB_TEST_HOST ?= db-test
DB_TEST_PORT ?= 5432
DB_TEST_USER ?= postgres
DB_TEST_PASSWORD ?= postgres
DB_TEST_NAME ?= gen_test
DB_TEST_SSLMODE ?= disable

# Construct database URL
DATABASE_URL = postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)
DATABASE_TEST_URL = postgres://$(DB_TEST_USER):$(DB_TEST_PASSWORD)@$(DB_TEST_HOST):$(DB_TEST_PORT)/$(DB_TEST_NAME)?sslmode=$(DB_TEST_SSLMODE)

.PHONY: migrate migrate-up migrate-down migrate-version migrate-create

# Run migration up
migrate:
	docker compose exec migration migrate -source file://./ -database '$(DATABASE_URL)' up
	docker compose exec migration migrate -source file://./ -database '$(DATABASE_TEST_URL)' up

# Alias for migrate
migrate-up: migrate

# Run migration down
migrate-down:
	docker compose exec migration migrate -source file://./ -database '$(DATABASE_URL)' down 1
	docker compose exec migration migrate -source file://./ -database '$(DATABASE_TEST_URL)' down 1

# Show current migration version
migrate-version:
	docker compose exec migration migrate -source file://./ -database '$(DATABASE_URL)' version
	docker compose exec migration migrate -source file://./ -database '$(DATABASE_TEST_URL)' version

# Create new migration file
migrate-create:
	@read -p "Enter migration name: " name; \
	docker compose exec migration migrate create -ext sql -dir ./ -seq $$name

.PHONY: logs
logs:
	docker logs gen-api -f --tail 100

.PHONY: generate-models
generate-models:
	@cd backend && rm -rf ./internal/domain/query/*.gen.go ./internal/domain/model/*.gen.go
	@cd backend && go run ./cmd/gormgen/generate_all/main.go

.PHONY: exec-schema
exec-schema: ## sqlファイルをコンテナに流す
	cat ./backend/migrations/*.up.sql > ./backend/migrations/schema.sql
	docker cp backend/migrations/schema.sql db:/schema.sql
	docker cp backend/migrations/schema.sql db-test:/schema.sql
	docker exec -it db psql -U postgres -d gen -f /schema.sql
	docker exec -it db-test psql -U postgres -d gen_test -f /schema.sql
	rm ./backend/migrations/schema.sql

.PHONY: swag
swag: ## swagger更新
	@docker compose exec gen-api swag init -g ./cmd/api/main.go
	@cd frontend && pnpm generate