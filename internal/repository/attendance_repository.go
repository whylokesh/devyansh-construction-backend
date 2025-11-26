package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/whylokesh/devyansh-construction-backend/internal/models"
)

type AttendanceRepository struct {
	db *pgxpool.Pool
}

func NewAttendanceRepository(db *pgxpool.Pool) *AttendanceRepository {
	return &AttendanceRepository{db: db}
}

func (r *AttendanceRepository) CreateAttendance(ctx context.Context, attendance *models.Attendance) error {
	query := `
		INSERT INTO attendance (worker_id, site_id, date, status, note)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at, updated_at
	`
	err := r.db.QueryRow(ctx, query,
		attendance.WorkerID,
		attendance.SiteID,
		attendance.Date,
		attendance.Status,
		attendance.Note,
	).Scan(&attendance.ID, &attendance.CreatedAt, &attendance.UpdatedAt)

	if err != nil {
		return fmt.Errorf("failed to create attendance: %w", err)
	}
	return nil
}

func (r *AttendanceRepository) GetAttendance(ctx context.Context, workerID, siteID int, date time.Time) (*models.Attendance, error) {
	query := `
		SELECT id, worker_id, site_id, date, status, note, created_at, updated_at
		FROM attendance
		WHERE worker_id = $1 AND site_id = $2 AND date = $3
	`
	var attendance models.Attendance
	err := r.db.QueryRow(ctx, query, workerID, siteID, date).Scan(
		&attendance.ID,
		&attendance.WorkerID,
		&attendance.SiteID,
		&attendance.Date,
		&attendance.Status,
		&attendance.Note,
		&attendance.CreatedAt,
		&attendance.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get attendance: %w", err)
	}
	return &attendance, nil
}

func (r *AttendanceRepository) GetAttendanceByID(ctx context.Context, id int) (*models.Attendance, error) {
	query := `
		SELECT id, worker_id, site_id, date, status, note, created_at, updated_at
		FROM attendance
		WHERE id = $1
	`
	var attendance models.Attendance
	err := r.db.QueryRow(ctx, query, id).Scan(
		&attendance.ID,
		&attendance.WorkerID,
		&attendance.SiteID,
		&attendance.Date,
		&attendance.Status,
		&attendance.Note,
		&attendance.CreatedAt,
		&attendance.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get attendance by id: %w", err)
	}
	return &attendance, nil
}

func (r *AttendanceRepository) UpdateAttendance(ctx context.Context, attendance *models.Attendance) error {
	query := `
		UPDATE attendance
		SET status = $1, note = $2, updated_at = NOW()
		WHERE id = $3
		RETURNING updated_at
	`
	err := r.db.QueryRow(ctx, query,
		attendance.Status,
		attendance.Note,
		attendance.ID,
	).Scan(&attendance.UpdatedAt)

	if err != nil {
		return fmt.Errorf("failed to update attendance: %w", err)
	}
	return nil
}

func (r *AttendanceRepository) DeleteAttendance(ctx context.Context, id int) error {
	query := `DELETE FROM attendance WHERE id = $1`
	_, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete attendance: %w", err)
	}
	return nil
}

func (r *AttendanceRepository) ListAttendanceBySite(ctx context.Context, siteID int, date time.Time) ([]models.Attendance, error) {
	query := `
		SELECT id, worker_id, site_id, date, status, note, created_at, updated_at
		FROM attendance
		WHERE site_id = $1 AND date = $2
	`
	rows, err := r.db.Query(ctx, query, siteID, date)
	if err != nil {
		return nil, fmt.Errorf("failed to list attendance by site: %w", err)
	}
	defer rows.Close()

	var attendances []models.Attendance
	for rows.Next() {
		var a models.Attendance
		if err := rows.Scan(
			&a.ID,
			&a.WorkerID,
			&a.SiteID,
			&a.Date,
			&a.Status,
			&a.Note,
			&a.CreatedAt,
			&a.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to scan attendance: %w", err)
		}
		attendances = append(attendances, a)
	}
	return attendances, nil
}

func (r *AttendanceRepository) ListAttendanceByWorker(ctx context.Context, workerID int, startDate, endDate time.Time) ([]models.Attendance, error) {
	query := `
		SELECT id, worker_id, site_id, date, status, note, created_at, updated_at
		FROM attendance
		WHERE worker_id = $1 AND date >= $2 AND date <= $3
		ORDER BY date
	`
	rows, err := r.db.Query(ctx, query, workerID, startDate, endDate)
	if err != nil {
		return nil, fmt.Errorf("failed to list attendance by worker: %w", err)
	}
	defer rows.Close()

	var attendances []models.Attendance
	for rows.Next() {
		var a models.Attendance
		if err := rows.Scan(
			&a.ID,
			&a.WorkerID,
			&a.SiteID,
			&a.Date,
			&a.Status,
			&a.Note,
			&a.CreatedAt,
			&a.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to scan attendance: %w", err)
		}
		attendances = append(attendances, a)
	}
	return attendances, nil
}
