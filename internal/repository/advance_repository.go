package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/whylokesh/devyansh-construction-backend/internal/models"
)

type AdvanceRepository struct {
	db *pgxpool.Pool
}

func NewAdvanceRepository(db *pgxpool.Pool) *AdvanceRepository {
	return &AdvanceRepository{db: db}
}

func (r *AdvanceRepository) CreateAdvance(ctx context.Context, advance *models.Advance) error {
	query := `
		INSERT INTO advances (worker_id, amount, given_on, note, created_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`
	err := r.db.QueryRow(ctx, query, advance.WorkerID, advance.Amount, advance.GivenOn, advance.Note, time.Now()).Scan(&advance.ID)
	if err != nil {
		return fmt.Errorf("failed to create advance: %w", err)
	}
	return nil
}

func (r *AdvanceRepository) GetAdvanceByID(ctx context.Context, id int) (*models.Advance, error) {
	query := `
		SELECT id, worker_id, amount, given_on, note, created_at
		FROM advances
		WHERE id = $1
	`
	var advance models.Advance
	err := r.db.QueryRow(ctx, query, id).Scan(&advance.ID, &advance.WorkerID, &advance.Amount, &advance.GivenOn, &advance.Note, &advance.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to get advance by id: %w", err)
	}
	return &advance, nil
}

func (r *AdvanceRepository) UpdateAdvance(ctx context.Context, advance *models.Advance) error {
	query := `
		UPDATE advances
		SET amount = $1, given_on = $2, note = $3
		WHERE id = $4
	`
	_, err := r.db.Exec(ctx, query, advance.Amount, advance.GivenOn, advance.Note, advance.ID)
	if err != nil {
		return fmt.Errorf("failed to update advance: %w", err)
	}
	return nil
}

func (r *AdvanceRepository) DeleteAdvance(ctx context.Context, id int) error {
	query := `DELETE FROM advances WHERE id = $1`
	_, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete advance: %w", err)
	}
	return nil
}

func (r *AdvanceRepository) ListAdvancesByWorker(ctx context.Context, workerID int) ([]models.Advance, error) {
	query := `
		SELECT id, worker_id, amount, given_on, note, created_at
		FROM advances
		WHERE worker_id = $1
		ORDER BY given_on DESC
	`
	rows, err := r.db.Query(ctx, query, workerID)
	if err != nil {
		return nil, fmt.Errorf("failed to list advances by worker: %w", err)
	}
	defer rows.Close()

	var advances []models.Advance
	for rows.Next() {
		var advance models.Advance
		if err := rows.Scan(&advance.ID, &advance.WorkerID, &advance.Amount, &advance.GivenOn, &advance.Note, &advance.CreatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan advance: %w", err)
		}
		advances = append(advances, advance)
	}
	return advances, nil
}

func (r *AdvanceRepository) ListAdvances(ctx context.Context) ([]models.Advance, error) {
	query := `
		SELECT id, worker_id, amount, given_on, note, created_at
		FROM advances
		ORDER BY given_on DESC
	`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to list advances: %w", err)
	}
	defer rows.Close()

	var advances []models.Advance
	for rows.Next() {
		var advance models.Advance
		if err := rows.Scan(&advance.ID, &advance.WorkerID, &advance.Amount, &advance.GivenOn, &advance.Note, &advance.CreatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan advance: %w", err)
		}
		advances = append(advances, advance)
	}
	return advances, nil
}
