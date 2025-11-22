package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/whylokesh/devyansh-construction-backend/internal/config"
	"github.com/whylokesh/devyansh-construction-backend/internal/db"
	"github.com/whylokesh/devyansh-construction-backend/internal/routes"
)

func main() {
	cfg := config.LoadConfig()
	db.InitDB(cfg)

	r := chi.NewRouter()
	routes.RegisterRoutes(r)

	http.ListenAndServe(":8080", r)
}
