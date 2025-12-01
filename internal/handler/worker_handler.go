package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/whylokesh/devyansh-construction-backend/internal/models"
	"github.com/whylokesh/devyansh-construction-backend/internal/service"
	"github.com/whylokesh/devyansh-construction-backend/internal/utils"
)

type WorkerHandler struct {
	service *service.WorkerService
}

func NewWorkerHandler(service *service.WorkerService) *WorkerHandler {
	return &WorkerHandler{service: service}
}

func (h *WorkerHandler) CreateWorker(w http.ResponseWriter, r *http.Request) {
	var worker models.Worker
	if err := json.NewDecoder(r.Body).Decode(&worker); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := h.service.CreateWorker(r.Context(), &worker); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, "Worker created successfully", worker)
}

func (h *WorkerHandler) GetWorkerByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid worker ID")
		return
	}

	worker, err := h.service.GetWorkerByID(r.Context(), id)
	if err != nil {
		utils.RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, "Worker retrieved successfully", worker)
}

func (h *WorkerHandler) UpdateWorker(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid worker ID")
		return
	}

	var worker models.Worker
	if err := json.NewDecoder(r.Body).Decode(&worker); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	worker.ID = id

	if err := h.service.UpdateWorker(r.Context(), &worker); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, "Worker updated successfully", worker)
}

func (h *WorkerHandler) DeleteWorker(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid worker ID")
		return
	}

	if err := h.service.DeleteWorker(r.Context(), id); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, "Worker deleted successfully", nil)
}

func (h *WorkerHandler) ListWorkers(w http.ResponseWriter, r *http.Request) {
	workers, err := h.service.ListWorkers(r.Context())
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, "Workers retrieved successfully", workers)
}
