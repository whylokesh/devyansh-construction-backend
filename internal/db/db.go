package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/whylokesh/devyansh-construction-backend/internal/config"
)

var DB *pgxpool.Pool

func InitDB(cfg *config.Config) {
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	if err = pool.Ping(ctx); err != nil {
		log.Fatalf("Unable to ping database: %v", err)
	}

	log.Println("Connected to PostgreSQL ✅")
	DB = pool
}