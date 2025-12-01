package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/whylokesh/devyansh-construction-backend/internal/handler"
	"github.com/whylokesh/devyansh-construction-backend/internal/middleware"
	"github.com/whylokesh/devyansh-construction-backend/internal/models"
)

func RegisterAdvanceRoutes(r chi.Router, advanceHandler *handler.AdvanceHandler, authMiddleware *middleware.AuthMiddleware) {
	r.Group(func(r chi.Router) {
		r.Use(authMiddleware.RequireRole(models.RoleAdmin))
		r.Post("/", advanceHandler.CreateAdvance)
		r.Get("/{id}", advanceHandler.GetAdvance)
		r.Put("/{id}", advanceHandler.UpdateAdvance)
		r.Delete("/{id}", advanceHandler.DeleteAdvance)
		r.Get("/", advanceHandler.ListAdvances)
	})
}
