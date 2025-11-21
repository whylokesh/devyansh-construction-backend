include .env
export

# 🧱 Migrations
migrate-up:
	migrate -path internal/db/migrations -database "$(DB_URL)" up

migrate-down:
	migrate -path internal/db/migrations -database "$(DB_URL)" down

create-migration:
	migrate create -ext sql -dir internal/db/migrations -seq $(name)

# 🚀 Run server
run:
	go run cmd/server/main.go