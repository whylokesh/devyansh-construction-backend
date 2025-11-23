package repository

import (
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

func TestNewRepositories(t *testing.T) {
	// This test just verifies that the constructors and structs are defined correctly.
	// It does not actually connect to a database.
	var db *pgxpool.Pool

	_ = NewUserRepository(db)
	_ = NewWorkerRepository(db)
	_ = NewSiteRepository(db)
	_ = NewAttendanceRepository(db)
	_ = NewFinanceRepository(db)
	_ = NewSnapshotRepository(db)
}
