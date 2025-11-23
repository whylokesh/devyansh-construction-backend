package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/whylokesh/devyansh-construction-backend/internal/config"
	"github.com/whylokesh/devyansh-construction-backend/internal/db"
	"github.com/whylokesh/devyansh-construction-backend/internal/handler"
	"github.com/whylokesh/devyansh-construction-backend/internal/repository"
	"github.com/whylokesh/devyansh-construction-backend/internal/routes"
	"github.com/whylokesh/devyansh-construction-backend/internal/service"
)

func main() {
	cfg := config.LoadConfig()
	db.InitDB(cfg)

	// Repositories
	userRepo := repository.NewUserRepository(db.DB)

	// Services
	userService := service.NewUserService(userRepo, cfg.JWTSecret)

	// Handlers
	userHandler := handler.NewUserHandler(userService)

	r := chi.NewRouter()
	routes.RegisterRoutes(r, userHandler)

	log.Println("Server starting on http://localhost:8080 🚀")
	http.ListenAndServe(":8080", r)
}
