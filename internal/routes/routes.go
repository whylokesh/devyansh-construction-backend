package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/whylokesh/devyansh-construction-backend/internal/handler"
)

func RegisterRoutes(r chi.Router, userHandler *handler.UserHandler) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("devyansh-construction-backend is workring"))
	})

	r.Route("/api", func(r chi.Router) {
		RegisterUserRoutes(r, userHandler)
	})
}