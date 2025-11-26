package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/whylokesh/devyansh-construction-backend/internal/handler"
)

func RegisterUserRoutes(r chi.Router, userHandler *handler.UserHandler) {
	// Routes for user
	r.Post("/signup", userHandler.Signup)
	r.Post("/login", userHandler.Login)
}
