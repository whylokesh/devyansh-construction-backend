package service

import (
	"context"
	"time"

	"github.com/whylokesh/devyansh-construction-backend/internal/models"
	"github.com/whylokesh/devyansh-construction-backend/internal/repository"
)

type AttendanceService struct {
	repo *repository.AttendanceRepository
}

func NewAttendanceService(repo *repository.AttendanceRepository) *AttendanceService {
	return &AttendanceService{repo: repo}
}

func (s *AttendanceService) CreateAttendance(ctx context.Context, attendance *models.Attendance) error {
	return s.repo.CreateAttendance(ctx, attendance)
}

func (s *AttendanceService) GetAttendanceByID(ctx context.Context, id int) (*models.Attendance, error) {
	return s.repo.GetAttendanceByID(ctx, id)
}

func (s *AttendanceService) UpdateAttendance(ctx context.Context, attendance *models.Attendance) error {
	return s.repo.UpdateAttendance(ctx, attendance)
}

func (s *AttendanceService) DeleteAttendance(ctx context.Context, id int) error {
	return s.repo.DeleteAttendance(ctx, id)
}

func (s *AttendanceService) ListAttendanceBySite(ctx context.Context, siteID int, date time.Time) ([]models.Attendance, error) {
	return s.repo.ListAttendanceBySite(ctx, siteID, date)
}

func (s *AttendanceService) ListAttendanceByWorker(ctx context.Context, workerID int, startDate, endDate time.Time) ([]models.Attendance, error) {
	return s.repo.ListAttendanceByWorker(ctx, workerID, startDate, endDate)
}
