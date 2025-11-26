package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/whylokesh/devyansh-construction-backend/internal/handler"
	"github.com/whylokesh/devyansh-construction-backend/internal/middleware"
	"github.com/whylokesh/devyansh-construction-backend/internal/models"
)

func RegisterSiteRoutes(r chi.Router, siteHandler *handler.SiteHandler, authMiddleware *middleware.AuthMiddleware) {
	// Public routes
	r.Get("/", siteHandler.ListSites)
	r.Get("/{id}", siteHandler.GetSiteByID)

	// Protected routes (Admin only)
	r.Group(func(r chi.Router) {
		r.Use(authMiddleware.RequireRole(models.RoleAdmin))
		r.Post("/", siteHandler.CreateSite)
		r.Put("/{id}", siteHandler.UpdateSite)
		r.Delete("/{id}", siteHandler.DeleteSite)
	})
}
