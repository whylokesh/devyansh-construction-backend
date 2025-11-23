package models

import (
	"encoding/json"
	"time"
)

type UserRole string

const (
	RoleAdmin      UserRole = "admin"
	RoleAccountant UserRole = "accountant"
)

type User struct {
	ID                int             `json:"id"`
	Name              string          `json:"name"`
	Email             string          `json:"email"`
	Role              UserRole        `json:"role"`
	PasswordHash      string          `json:"-"`
	AdditionalDetails json.RawMessage `json:"additional_details"`
	CreatedAt         time.Time       `json:"created_at"`
	UpdatedAt         time.Time       `json:"updated_at"`
}
