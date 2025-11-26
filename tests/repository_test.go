package tests

import (
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/whylokesh/devyansh-construction-backend/internal/repository"
)

func TestNewRepositories(t *testing.T) {
	// This test just verifies that the constructors and structs are defined correctly.
	// It does not actually connect to a database.
	var db *pgxpool.Pool

	_ = repository.NewUserRepository(db)
	_ = repository.NewWorkerRepository(db)
	_ = repository.NewSiteRepository(db)
	_ = repository.NewAttendanceRepository(db)
	_ = repository.NewFinanceRepository(db)
	_ = repository.NewSnapshotRepository(db)
}
