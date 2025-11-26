package service

import (
	"context"

	"github.com/whylokesh/devyansh-construction-backend/internal/models"
	"github.com/whylokesh/devyansh-construction-backend/internal/repository"
)

type SiteService struct {
	repo *repository.SiteRepository
}

func NewSiteService(repo *repository.SiteRepository) *SiteService {
	return &SiteService{repo: repo}
}

func (s *SiteService) CreateSite(ctx context.Context, site *models.Site) error {
	return s.repo.CreateSite(ctx, site)
}

func (s *SiteService) GetSiteByID(ctx context.Context, id int) (*models.Site, error) {
	return s.repo.GetSiteByID(ctx, id)
}

func (s *SiteService) UpdateSite(ctx context.Context, site *models.Site) error {
	return s.repo.UpdateSite(ctx, site)
}

func (s *SiteService) DeleteSite(ctx context.Context, id int) error {
	return s.repo.DeleteSite(ctx, id)
}

func (s *SiteService) ListSites(ctx context.Context) ([]models.Site, error) {
	return s.repo.ListSites(ctx)
}
