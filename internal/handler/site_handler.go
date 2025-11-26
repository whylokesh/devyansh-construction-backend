package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/whylokesh/devyansh-construction-backend/internal/models"
	"github.com/whylokesh/devyansh-construction-backend/internal/service"
)

type SiteHandler struct {
	service *service.SiteService
}

func NewSiteHandler(service *service.SiteService) *SiteHandler {
	return &SiteHandler{service: service}
}

func (h *SiteHandler) CreateSite(w http.ResponseWriter, r *http.Request) {
	var site models.Site
	if err := json.NewDecoder(r.Body).Decode(&site); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := h.service.CreateSite(r.Context(), &site); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusCreated, "Site created successfully", site)
}

func (h *SiteHandler) GetSiteByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid site ID")
		return
	}

	site, err := h.service.GetSiteByID(r.Context(), id)
	if err != nil {
		RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, "Site retrieved successfully", site)
}

func (h *SiteHandler) UpdateSite(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid site ID")
		return
	}

	var site models.Site
	if err := json.NewDecoder(r.Body).Decode(&site); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	site.ID = id

	if err := h.service.UpdateSite(r.Context(), &site); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, "Site updated successfully", site)
}

func (h *SiteHandler) DeleteSite(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid site ID")
		return
	}

	if err := h.service.DeleteSite(r.Context(), id); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, "Site deleted successfully", nil)
}

func (h *SiteHandler) ListSites(w http.ResponseWriter, r *http.Request) {
	sites, err := h.service.ListSites(r.Context())
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, "Sites retrieved successfully", sites)
}
