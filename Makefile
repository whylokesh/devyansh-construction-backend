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
	migrate -path internal/db/migrations -database "postgres://postgres:H32sOM3w8ehg3picaJwDFewFRnF8gl6IBLsOyZAwqgytor9Zhq0dD5KZgwkjAlak@31.97.205.243:5432/postgres?sslmode=disable" up

migrate-prod-down:
	migrate -path internal/db/migrations -database "postgres://postgres:H32sOM3w8ehg3picaJwDFewFRnF8gl6IBLsOyZAwqgytor9Zhq0dD5KZgwkjAlak@31.97.205.243:5432/postgres?sslmode=disable" down

# 🚀 Run server
run:
	go run cmd/server/main.go
