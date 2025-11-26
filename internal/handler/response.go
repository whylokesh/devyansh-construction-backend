package handler

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func RespondWithJSON(w http.ResponseWriter, status int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(Response{
		Message: message,
		Data:    data,
	})
}

func RespondWithError(w http.ResponseWriter, status int, message string) {
	RespondWithJSON(w, status, message, nil)
}
