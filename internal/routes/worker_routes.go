package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/whylokesh/devyansh-construction-backend/internal/handler"
	"github.com/whylokesh/devyansh-construction-backend/internal/middleware"
	"github.com/whylokesh/devyansh-construction-backend/internal/models"
)

func RegisterWorkerRoutes(r chi.Router, workerHandler *handler.WorkerHandler, authMiddleware *middleware.AuthMiddleware) {
	// Public routes
	r.Get("/", workerHandler.ListWorkers)
	r.Get("/{id}", workerHandler.GetWorkerByID)

	// Protected routes (Admin only)
	r.Group(func(r chi.Router) {
		r.Use(authMiddleware.RequireRole(models.RoleAdmin))
		r.Post("/", workerHandler.CreateWorker)
		r.Put("/{id}", workerHandler.UpdateWorker)
		r.Delete("/{id}", workerHandler.DeleteWorker)
	})
}
