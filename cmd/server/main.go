package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/whylokesh/devyansh-construction-backend/internal/config"
	"github.com/whylokesh/devyansh-construction-backend/internal/db"
	"github.com/whylokesh/devyansh-construction-backend/internal/repository"
	"github.com/whylokesh/devyansh-construction-backend/internal/routes"
)

func main() {
	cfg := config.LoadConfig()
	db.InitDB(cfg)

	userRepo := repository.NewUserRepository(db.DB)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	r := chi.NewRouter()
	routes.RegisterRoutes(r, userHandler)

	http.ListenAndServe(":8080", r)
}
