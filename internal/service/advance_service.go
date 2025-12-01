package service

import (
	"context"

	"github.com/whylokesh/devyansh-construction-backend/internal/models"
	"github.com/whylokesh/devyansh-construction-backend/internal/repository"
)

type AdvanceService struct {
	repo *repository.AdvanceRepository
}

func NewAdvanceService(repo *repository.AdvanceRepository) *AdvanceService {
	return &AdvanceService{repo: repo}
}

func (s *AdvanceService) CreateAdvance(ctx context.Context, advance *models.Advance) error {
	return s.repo.CreateAdvance(ctx, advance)
}

func (s *AdvanceService) GetAdvanceByID(ctx context.Context, id int) (*models.Advance, error) {
	return s.repo.GetAdvanceByID(ctx, id)
}

func (s *AdvanceService) UpdateAdvance(ctx context.Context, advance *models.Advance) error {
	return s.repo.UpdateAdvance(ctx, advance)
}

func (s *AdvanceService) DeleteAdvance(ctx context.Context, id int) error {
	return s.repo.DeleteAdvance(ctx, id)
}

func (s *AdvanceService) ListAdvancesByWorker(ctx context.Context, workerID int) ([]models.Advance, error) {
	return s.repo.ListAdvancesByWorker(ctx, workerID)
}

func (s *AdvanceService) ListAdvances(ctx context.Context) ([]models.Advance, error) {
	return s.repo.ListAdvances(ctx)
}
