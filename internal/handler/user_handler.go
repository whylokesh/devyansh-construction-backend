package handler

import (
	"encoding/json"
	"net/http"

	"github.com/whylokesh/devyansh-construction-backend/internal/service"
	"github.com/whylokesh/devyansh-construction-backend/internal/utils"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) Signup(w http.ResponseWriter, r *http.Request) {
	var input service.RegisterUserInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid input")
		return
	}

	user, err := h.service.RegisterUser(r.Context(), input)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, "User registered successfully", user)
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid input")
		return
	}

	token, user, err := h.service.LoginUser(r.Context(), input.Email, input.Password)
	if err != nil {
		utils.RespondWithError(w, http.StatusUnauthorized, "Invalid email or password")
		return
	}

	data := map[string]interface{}{
		"token": token,
		"user":  user,
	}

	utils.RespondWithJSON(w, http.StatusOK, "Login successful", data)
}
