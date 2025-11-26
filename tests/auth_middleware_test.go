package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/whylokesh/devyansh-construction-backend/internal/middleware"
	"github.com/whylokesh/devyansh-construction-backend/internal/models"
)

func TestAuthMiddleware_RequireRole(t *testing.T) {
	secret := "testsecret"
	authMiddleware := middleware.NewAuthMiddleware(secret)

	tests := []struct {
		name           string
		roles          []models.UserRole
		tokenRole      models.UserRole
		expectedStatus int
		setupToken     func() string
	}{
		{
			name:           "Valid Token and Role",
			roles:          []models.UserRole{models.RoleAdmin},
			tokenRole:      models.RoleAdmin,
			expectedStatus: http.StatusOK,
			setupToken: func() string {
				token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
					"user_id": 1,
					"role":    models.RoleAdmin,
					"exp":     time.Now().Add(time.Hour).Unix(),
				})
				tokenString, _ := token.SignedString([]byte(secret))
				return "Bearer " + tokenString
			},
		},
		{
			name:           "Valid Token but Invalid Role",
			roles:          []models.UserRole{models.RoleAdmin},
			tokenRole:      models.RoleAccountant,
			expectedStatus: http.StatusForbidden,
			setupToken: func() string {
				token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
					"user_id": 2,
					"role":    models.RoleAccountant,
					"exp":     time.Now().Add(time.Hour).Unix(),
				})
				tokenString, _ := token.SignedString([]byte(secret))
				return "Bearer " + tokenString
			},
		},
		{
			name:           "Missing Token",
			roles:          []models.UserRole{models.RoleAdmin},
			expectedStatus: http.StatusUnauthorized,
			setupToken: func() string {
				return ""
			},
		},
		{
			name:           "Invalid Token Format",
			roles:          []models.UserRole{models.RoleAdmin},
			expectedStatus: http.StatusUnauthorized,
			setupToken: func() string {
				return "InvalidToken"
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := authMiddleware.RequireRole(tt.roles...)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			}))

			req := httptest.NewRequest("GET", "/", nil)
			if tt.setupToken != nil {
				token := tt.setupToken()
				if token != "" {
					req.Header.Set("Authorization", token)
				}
			}

			rr := httptest.NewRecorder()
			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.expectedStatus)
			}
		})
	}
}
