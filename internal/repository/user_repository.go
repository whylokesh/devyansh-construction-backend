package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/whylokesh/devyansh-construction-backend/internal/models"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *models.User) error {
	query := `
		INSERT INTO users (name, email, role, password_hash, additional_details)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at, updated_at
	`
	err := r.db.QueryRow(ctx, query,
		user.Name,
		user.Email,
		user.Role,
		user.PasswordHash,
		user.AdditionalDetails,
	).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

func (r *UserRepository) GetUserByID(ctx context.Context, id int) (*models.User, error) {
	query := `
		SELECT id, name, email, role, password_hash, additional_details, created_at, updated_at
		FROM users
		WHERE id = $1
	`
	var user models.User
	err := r.db.QueryRow(ctx, query, id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Role,
		&user.PasswordHash,
		&user.AdditionalDetails,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by id: %w", err)
	}
	return &user, nil
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	query := `
		SELECT id, name, email, role, password_hash, additional_details, created_at, updated_at
		FROM users
		WHERE email = $1
	`
	var user models.User
	err := r.db.QueryRow(ctx, query, email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Role,
		&user.PasswordHash,
		&user.AdditionalDetails,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by email: %w", err)
	}
	return &user, nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, user *models.User) error {
	query := `
		UPDATE users
		SET name = $1, email = $2, role = $3, password_hash = $4, additional_details = $5, updated_at = NOW()
		WHERE id = $6
		RETURNING updated_at
	`
	err := r.db.QueryRow(ctx, query,
		user.Name,
		user.Email,
		user.Role,
		user.PasswordHash,
		user.AdditionalDetails,
		user.ID,
	).Scan(&user.UpdatedAt)

	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}
	return nil
}

func (r *UserRepository) DeleteUser(ctx context.Context, id int) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}
