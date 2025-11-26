package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/whylokesh/devyansh-construction-backend/internal/config"
	"github.com/whylokesh/devyansh-construction-backend/internal/db"
	"github.com/whylokesh/devyansh-construction-backend/internal/handler"
	"github.com/whylokesh/devyansh-construction-backend/internal/middleware"
	"github.com/whylokesh/devyansh-construction-backend/internal/repository"
	"github.com/whylokesh/devyansh-construction-backend/internal/routes"
	"github.com/whylokesh/devyansh-construction-backend/internal/service"
)

func main() {
	cfg := config.LoadConfig()
	db.InitDB(cfg)

	// Repositories
	userRepo := repository.NewUserRepository(db.DB)
	siteRepo := repository.NewSiteRepository(db.DB)
	workerRepo := repository.NewWorkerRepository(db.DB)

	// Services
	userService := service.NewUserService(userRepo, cfg.JWTSecret)
	siteService := service.NewSiteService(siteRepo)
	workerService := service.NewWorkerService(workerRepo)

	// Handlers
	userHandler := handler.NewUserHandler(userService)
	siteHandler := handler.NewSiteHandler(siteService)
	workerHandler := handler.NewWorkerHandler(workerService)

	// Middleware
	authMiddleware := middleware.NewAuthMiddleware(cfg.JWTSecret)

	r := chi.NewRouter()
	routes.RegisterRoutes(r, userHandler, siteHandler, workerHandler, authMiddleware)

	log.Println("Server starting on http://localhost:8080 🚀")
	http.ListenAndServe(":8080", r)
}
