include .env
export

# 🧱 Migrations (local)
migrate-up:
	migrate -path internal/db/migrations -database "$(DB_URL)" up

migrate-down:
	migrate -path internal/db/migrations -database "$(DB_URL)" down

create-migration:
	migrate create -ext sql -dir internal/db/migrations -seq $(name)

# 🧱 Migrations (production)
migrate-prod-up:
	migrate -path internal/db/migrations -database "$(CLOUD_DB_URL_EXTERNAL)" up

migrate-prod-down:
	migrate -path internal/db/migrations -database "$(CLOUD_DB_URL_EXTERNAL)" down

# 🚀 Run server
run:
	go run cmd/server/main.go

# 🔄 Live Reload
watch:
	air
