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

type AdvanceHandler struct {
	service *service.AdvanceService
}

func NewAdvanceHandler(service *service.AdvanceService) *AdvanceHandler {
	return &AdvanceHandler{service: service}				
}

func (h *AdvanceHandler) CreateAdvance(w http.ResponseWriter, r *http.Request) {
	var advance models.Advance
	if err := json.NewDecoder(r.Body).Decode(&advance); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := h.service.CreateAdvance(r.Context(), &advance); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, "Advance created successfully", advance)
}

func (h *AdvanceHandler) GetAdvance(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	advance, err := h.service.GetAdvanceByID(r.Context(), id)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, "Advance retrieved successfully", advance)
}

func (h *AdvanceHandler) UpdateAdvance(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	var advance models.Advance
	if err := json.NewDecoder(r.Body).Decode(&advance); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	advance.ID = id

	if err := h.service.UpdateAdvance(r.Context(), &advance); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, "Advance updated successfully", advance)
}

func (h *AdvanceHandler) DeleteAdvance(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	if err := h.service.DeleteAdvance(r.Context(), id); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, "Advance deleted successfully", nil)
}

func (h *AdvanceHandler) ListAdvances(w http.ResponseWriter, r *http.Request) {
	workerIDStr := r.URL.Query().Get("worker_id")
	var advances []models.Advance
	var err error

	if workerIDStr != "" {
		var workerID int
		workerID, err = strconv.Atoi(workerIDStr)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid worker_id")
			return
		}
		advances, err = h.service.ListAdvancesByWorker(r.Context(), workerID)
	} else {
		advances, err = h.service.ListAdvances(r.Context())
	}

	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, "Advances retrieved successfully", advances)
}
