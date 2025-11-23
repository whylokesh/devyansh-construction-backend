package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/whylokesh/devyansh-construction-backend/internal/models"
)

type SiteRepository struct {
	db *pgxpool.Pool
}

func NewSiteRepository(db *pgxpool.Pool) *SiteRepository {
	return &SiteRepository{db: db}
}

func (r *SiteRepository) CreateSite(ctx context.Context, site *models.Site) error {
	query := `
		INSERT INTO sites (name, client_name, client_phone, location, site_documents, additional_details, start_date, end_date, status)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id, created_at, updated_at
	`
	err := r.db.QueryRow(ctx, query,
		site.Name,
		site.ClientName,
		site.ClientPhone,
		site.Location,
		site.SiteDocuments,
		site.AdditionalDetails,
		site.StartDate,
		site.EndDate,
		site.Status,
	).Scan(&site.ID, &site.CreatedAt, &site.UpdatedAt)

	if err != nil {
		return fmt.Errorf("failed to create site: %w", err)
	}
	return nil
}

func (r *SiteRepository) GetSiteByID(ctx context.Context, id int) (*models.Site, error) {
	query := `
		SELECT id, name, client_name, client_phone, location, site_documents, additional_details, start_date, end_date, status, created_at, updated_at
		FROM sites
		WHERE id = $1
	`
	var site models.Site
	err := r.db.QueryRow(ctx, query, id).Scan(
		&site.ID,
		&site.Name,
		&site.ClientName,
		&site.ClientPhone,
		&site.Location,
		&site.SiteDocuments,
		&site.AdditionalDetails,
		&site.StartDate,
		&site.EndDate,
		&site.Status,
		&site.CreatedAt,
		&site.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get site by id: %w", err)
	}
	return &site, nil
}

func (r *SiteRepository) UpdateSite(ctx context.Context, site *models.Site) error {
	query := `
		UPDATE sites
		SET name = $1, client_name = $2, client_phone = $3, location = $4, site_documents = $5, additional_details = $6, start_date = $7, end_date = $8, status = $9, updated_at = NOW()
		WHERE id = $10
		RETURNING updated_at
	`
	err := r.db.QueryRow(ctx, query,
		site.Name,
		site.ClientName,
		site.ClientPhone,
		site.Location,
		site.SiteDocuments,
		site.AdditionalDetails,
		site.StartDate,
		site.EndDate,
		site.Status,
		site.ID,
	).Scan(&site.UpdatedAt)

	if err != nil {
		return fmt.Errorf("failed to update site: %w", err)
	}
	return nil
}

func (r *SiteRepository) DeleteSite(ctx context.Context, id int) error {
	query := `DELETE FROM sites WHERE id = $1`
	_, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete site: %w", err)
	}
	return nil
}

func (r *SiteRepository) ListSites(ctx context.Context) ([]models.Site, error) {
	query := `
		SELECT id, name, client_name, client_phone, location, site_documents, additional_details, start_date, end_date, status, created_at, updated_at
		FROM sites
		ORDER BY name
	`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to list sites: %w", err)
	}
	defer rows.Close()

	var sites []models.Site
	for rows.Next() {
		var s models.Site
		if err := rows.Scan(
			&s.ID,
			&s.Name,
			&s.ClientName,
			&s.ClientPhone,
			&s.Location,
			&s.SiteDocuments,
			&s.AdditionalDetails,
			&s.StartDate,
			&s.EndDate,
			&s.Status,
			&s.CreatedAt,
			&s.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to scan site: %w", err)
		}
		sites = append(sites, s)
	}
	return sites, nil
}
