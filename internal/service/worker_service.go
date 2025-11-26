package service

import (
	"context"

	"github.com/whylokesh/devyansh-construction-backend/internal/models"
	"github.com/whylokesh/devyansh-construction-backend/internal/repository"
)

type WorkerService struct {
	repo *repository.WorkerRepository
}

func NewWorkerService(repo *repository.WorkerRepository) *WorkerService {
	return &WorkerService{repo: repo}
}

func (s *WorkerService) CreateWorker(ctx context.Context, worker *models.Worker) error {
	return s.repo.CreateWorker(ctx, worker)
}

func (s *WorkerService) GetWorkerByID(ctx context.Context, id int) (*models.Worker, error) {
	return s.repo.GetWorkerByID(ctx, id)
}

func (s *WorkerService) UpdateWorker(ctx context.Context, worker *models.Worker) error {
	return s.repo.UpdateWorker(ctx, worker)
}

func (s *WorkerService) DeleteWorker(ctx context.Context, id int) error {
	return s.repo.DeleteWorker(ctx, id)
}

func (s *WorkerService) ListWorkers(ctx context.Context) ([]models.Worker, error) {
	return s.repo.ListWorkers(ctx)
}
