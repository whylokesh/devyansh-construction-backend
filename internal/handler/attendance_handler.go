package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/whylokesh/devyansh-construction-backend/internal/models"
	"github.com/whylokesh/devyansh-construction-backend/internal/service"
	"github.com/whylokesh/devyansh-construction-backend/internal/utils"
)

type AttendanceHandler struct {
	service *service.AttendanceService
}

func NewAttendanceHandler(service *service.AttendanceService) *AttendanceHandler {
	return &AttendanceHandler{service: service}
}

func (h *AttendanceHandler) CreateAttendance(w http.ResponseWriter, r *http.Request) {
	var attendance models.Attendance
	if err := json.NewDecoder(r.Body).Decode(&attendance); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := h.service.CreateAttendance(r.Context(), &attendance); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, "Attendance recorded successfully", attendance)
}

func (h *AttendanceHandler) GetAttendanceByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid attendance ID")
		return
	}

	attendance, err := h.service.GetAttendanceByID(r.Context(), id)
	if err != nil {
		utils.RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, "Attendance retrieved successfully", attendance)
}

func (h *AttendanceHandler) UpdateAttendance(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid attendance ID")
		return
	}

	var attendance models.Attendance
	if err := json.NewDecoder(r.Body).Decode(&attendance); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	attendance.ID = id

	if err := h.service.UpdateAttendance(r.Context(), &attendance); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, "Attendance updated successfully", attendance)
}

func (h *AttendanceHandler) DeleteAttendance(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid attendance ID")
		return
	}

	if err := h.service.DeleteAttendance(r.Context(), id); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, "Attendance deleted successfully", nil)
}

func (h *AttendanceHandler) ListAttendanceBySite(w http.ResponseWriter, r *http.Request) {
	siteIDStr := chi.URLParam(r, "siteId")
	siteID, err := strconv.Atoi(siteIDStr)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid site ID")
		return
	}

	dateStr := r.URL.Query().Get("date")
	var date time.Time
	if dateStr != "" {
		date, err = time.Parse("2006-01-02", dateStr)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid date format (YYYY-MM-DD)")
			return
		}
	} else {
		date = time.Now() // Default to today if not specified
	}

	attendances, err := h.service.ListAttendanceBySite(r.Context(), siteID, date)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, "Attendance list retrieved successfully", attendances)
}

func (h *AttendanceHandler) ListAttendanceByWorker(w http.ResponseWriter, r *http.Request) {
	workerIDStr := chi.URLParam(r, "workerId")
	workerID, err := strconv.Atoi(workerIDStr)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid worker ID")
		return
	}

	startDateStr := r.URL.Query().Get("startDate")
	endDateStr := r.URL.Query().Get("endDate")

	var startDate, endDate time.Time
	if startDateStr != "" {
		startDate, err = time.Parse("2006-01-02", startDateStr)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid start date format (YYYY-MM-DD)")
			return
		}
	} else {
		// Default to first day of current month
		now := time.Now()
		startDate = time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	}

	if endDateStr != "" {
		endDate, err = time.Parse("2006-01-02", endDateStr)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid end date format (YYYY-MM-DD)")
			return
		}
	} else {
		endDate = time.Now() // Default to today
	}

	attendances, err := h.service.ListAttendanceByWorker(r.Context(), workerID, startDate, endDate)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, "Attendance list retrieved successfully", attendances)
}
