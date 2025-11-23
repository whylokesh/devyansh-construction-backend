package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/whylokesh/devyansh-construction-backend/internal/models"
)

type SnapshotRepository struct {
	db *pgxpool.Pool
}

func NewSnapshotRepository(db *pgxpool.Pool) *SnapshotRepository {
	return &SnapshotRepository{db: db}
}

func (r *SnapshotRepository) CreateSnapshot(ctx context.Context, snapshot *models.SiteSummarySnapshot) error {
	query := `
		INSERT INTO site_summary_snapshots (site_id, summary_json)
		VALUES ($1, $2)
		RETURNING id, created_at
	`
	err := r.db.QueryRow(ctx, query,
		snapshot.SiteID,
		snapshot.SummaryJSON,
	).Scan(&snapshot.ID, &snapshot.CreatedAt)

	if err != nil {
		return fmt.Errorf("failed to create snapshot: %w", err)
	}
	return nil
}

func (r *SnapshotRepository) GetSnapshotByID(ctx context.Context, id int) (*models.SiteSummarySnapshot, error) {
	query := `
		SELECT id, site_id, summary_json, created_at
		FROM site_summary_snapshots
		WHERE id = $1
	`
	var snapshot models.SiteSummarySnapshot
	err := r.db.QueryRow(ctx, query, id).Scan(
		&snapshot.ID,
		&snapshot.SiteID,
		&snapshot.SummaryJSON,
		&snapshot.CreatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get snapshot by id: %w", err)
	}
	return &snapshot, nil
}
