package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/whylokesh/devyansh-construction-backend/internal/models"
)

type WorkerRepository struct {
	db *pgxpool.Pool
}

func NewWorkerRepository(db *pgxpool.Pool) *WorkerRepository {
	return &WorkerRepository{db: db}
}

func (r *WorkerRepository) CreateWorker(ctx context.Context, worker *models.Worker) error {
	query := `
		INSERT INTO workers (name, phone, skill, bill_rate, payout_rate, active_status, notes, additional_details)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id, created_at, updated_at
	`
	err := r.db.QueryRow(ctx, query,
		worker.Name,
		worker.Phone,
		worker.Skill,
		worker.BillRate,
		worker.PayoutRate,
		worker.ActiveStatus,
		worker.Notes,
		worker.AdditionalDetails,
	).Scan(&worker.ID, &worker.CreatedAt, &worker.UpdatedAt)

	if err != nil {
		return fmt.Errorf("failed to create worker: %w", err)
	}
	return nil
}

func (r *WorkerRepository) GetWorkerByID(ctx context.Context, id int) (*models.Worker, error) {
	query := `
		SELECT id, name, phone, skill, bill_rate, payout_rate, active_status, notes, additional_details, created_at, updated_at
		FROM workers
		WHERE id = $1
	`
	var worker models.Worker
	err := r.db.QueryRow(ctx, query, id).Scan(
		&worker.ID,
		&worker.Name,
		&worker.Phone,
		&worker.Skill,
		&worker.BillRate,
		&worker.PayoutRate,
		&worker.ActiveStatus,
		&worker.Notes,
		&worker.AdditionalDetails,
		&worker.CreatedAt,
		&worker.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get worker by id: %w", err)
	}
	return &worker, nil
}

func (r *WorkerRepository) UpdateWorker(ctx context.Context, worker *models.Worker) error {
	query := `
		UPDATE workers
		SET name = $1, phone = $2, skill = $3, bill_rate = $4, payout_rate = $5, active_status = $6, notes = $7, additional_details = $8, updated_at = NOW()
		WHERE id = $9
		RETURNING updated_at
	`
	err := r.db.QueryRow(ctx, query,
		worker.Name,
		worker.Phone,
		worker.Skill,
		worker.BillRate,
		worker.PayoutRate,
		worker.ActiveStatus,
		worker.Notes,
		worker.AdditionalDetails,
		worker.ID,
	).Scan(&worker.UpdatedAt)

	if err != nil {
		return fmt.Errorf("failed to update worker: %w", err)
	}
	return nil
}

func (r *WorkerRepository) DeleteWorker(ctx context.Context, id int) error {
	query := `DELETE FROM workers WHERE id = $1`
	_, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete worker: %w", err)
	}
	return nil
}

func (r *WorkerRepository) ListWorkers(ctx context.Context) ([]models.Worker, error) {
	query := `
		SELECT id, name, phone, skill, bill_rate, payout_rate, active_status, notes, additional_details, created_at, updated_at
		FROM workers
		ORDER BY name
	`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to list workers: %w", err)
	}
	defer rows.Close()

	var workers []models.Worker
	for rows.Next() {
		var w models.Worker
		if err := rows.Scan(
			&w.ID,
			&w.Name,
			&w.Phone,
			&w.Skill,
			&w.BillRate,
			&w.PayoutRate,
			&w.ActiveStatus,
			&w.Notes,
			&w.AdditionalDetails,
			&w.CreatedAt,
			&w.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to scan worker: %w", err)
		}
		workers = append(workers, w)
	}
	return workers, nil
}
