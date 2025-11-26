package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/whylokesh/devyansh-construction-backend/internal/handler"
	"github.com/whylokesh/devyansh-construction-backend/internal/middleware"
	"github.com/whylokesh/devyansh-construction-backend/internal/models"
)

func RegisterAttendanceRoutes(r chi.Router, attendanceHandler *handler.AttendanceHandler, authMiddleware *middleware.AuthMiddleware) {
	// Protected routes (Admin only for now, can be adjusted)
	r.Group(func(r chi.Router) {
		r.Use(authMiddleware.RequireRole(models.RoleAdmin))
		r.Post("/", attendanceHandler.CreateAttendance)
		r.Get("/{id}", attendanceHandler.GetAttendanceByID)
		r.Put("/{id}", attendanceHandler.UpdateAttendance)
		r.Delete("/{id}", attendanceHandler.DeleteAttendance)
		r.Get("/site/{siteId}", attendanceHandler.ListAttendanceBySite)
		r.Get("/worker/{workerId}", attendanceHandler.ListAttendanceByWorker)
	})
}
