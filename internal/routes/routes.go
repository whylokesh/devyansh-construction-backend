package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/whylokesh/devyansh-construction-backend/internal/handler"
	"github.com/whylokesh/devyansh-construction-backend/internal/middleware"
)

func RegisterRoutes(r chi.Router, userHandler *handler.UserHandler, siteHandler *handler.SiteHandler, workerHandler *handler.WorkerHandler, attendanceHandler *handler.AttendanceHandler, advanceHandler *handler.AdvanceHandler, authMiddleware *middleware.AuthMiddleware) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("devyansh-construction-backend is workring"))
	})

	r.Route("/api", func(r chi.Router) {
		r.Route("/users", func(r chi.Router) {
			RegisterUserRoutes(r, userHandler)
		})
		r.Route("/sites", func(r chi.Router) {
			RegisterSiteRoutes(r, siteHandler, authMiddleware)
		})
		r.Route("/workers", func(r chi.Router) {
			RegisterWorkerRoutes(r, workerHandler, authMiddleware)
		})
		r.Route("/attendance", func(r chi.Router) {
			RegisterAttendanceRoutes(r, attendanceHandler, authMiddleware)
		})
		r.Route("/advances", func(r chi.Router) {
			RegisterAdvanceRoutes(r, advanceHandler, authMiddleware)
		})
	})
}