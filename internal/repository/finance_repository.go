package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/whylokesh/devyansh-construction-backend/internal/models"
)

type FinanceRepository struct {
	db *pgxpool.Pool
}

func NewFinanceRepository(db *pgxpool.Pool) *FinanceRepository {
	return &FinanceRepository{db: db}
}

// Advances

func (r *FinanceRepository) CreateAdvance(ctx context.Context, advance *models.Advance) error {
	query := `
		INSERT INTO advances (worker_id, amount, given_on, note)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at
	`
	err := r.db.QueryRow(ctx, query,
		advance.WorkerID,
		advance.Amount,
		advance.GivenOn,
		advance.Note,
	).Scan(&advance.ID, &advance.CreatedAt)

	if err != nil {
		return fmt.Errorf("failed to create advance: %w", err)
	}
	return nil
}

func (r *FinanceRepository) GetAdvanceByID(ctx context.Context, id int) (*models.Advance, error) {
	query := `
		SELECT id, worker_id, amount, given_on, note, created_at
		FROM advances
		WHERE id = $1
	`
	var advance models.Advance
	err := r.db.QueryRow(ctx, query, id).Scan(
		&advance.ID,
		&advance.WorkerID,
		&advance.Amount,
		&advance.GivenOn,
		&advance.Note,
		&advance.CreatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get advance by id: %w", err)
	}
	return &advance, nil
}

// Payouts

func (r *FinanceRepository) CreatePayout(ctx context.Context, payout *models.Payout) error {
	query := `
		INSERT INTO payouts (period_start, period_end, payout_json, status)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at, updated_at
	`
	err := r.db.QueryRow(ctx, query,
		payout.PeriodStart,
		payout.PeriodEnd,
		payout.PayoutJSON,
		payout.Status,
	).Scan(&payout.ID, &payout.CreatedAt, &payout.UpdatedAt)

	if err != nil {
		return fmt.Errorf("failed to create payout: %w", err)
	}
	return nil
}

func (r *FinanceRepository) GetPayoutByID(ctx context.Context, id int) (*models.Payout, error) {
	query := `
		SELECT id, period_start, period_end, payout_json, status, created_at, updated_at
		FROM payouts
		WHERE id = $1
	`
	var payout models.Payout
	err := r.db.QueryRow(ctx, query, id).Scan(
		&payout.ID,
		&payout.PeriodStart,
		&payout.PeriodEnd,
		&payout.PayoutJSON,
		&payout.Status,
		&payout.CreatedAt,
		&payout.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get payout by id: %w", err)
	}
	return &payout, nil
}

// Bills

func (r *FinanceRepository) CreateBill(ctx context.Context, bill *models.Bill) error {
	query := `
		INSERT INTO bills (site_id, period_start, period_end, bill_json, status)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at, updated_at
	`
	err := r.db.QueryRow(ctx, query,
		bill.SiteID,
		bill.PeriodStart,
		bill.PeriodEnd,
		bill.BillJSON,
		bill.Status,
	).Scan(&bill.ID, &bill.CreatedAt, &bill.UpdatedAt)

	if err != nil {
		return fmt.Errorf("failed to create bill: %w", err)
	}
	return nil
}

func (r *FinanceRepository) GetBillByID(ctx context.Context, id int) (*models.Bill, error) {
	query := `
		SELECT id, site_id, period_start, period_end, bill_json, status, created_at, updated_at
		FROM bills
		WHERE id = $1
	`
	var bill models.Bill
	err := r.db.QueryRow(ctx, query, id).Scan(
		&bill.ID,
		&bill.SiteID,
		&bill.PeriodStart,
		&bill.PeriodEnd,
		&bill.BillJSON,
		&bill.Status,
		&bill.CreatedAt,
		&bill.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get bill by id: %w", err)
	}
	return &bill, nil
}
