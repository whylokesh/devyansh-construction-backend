package models

import (
	"encoding/json"
	"time"
)

type Worker struct {
	ID                int             `json:"id"`
	Name              string          `json:"name"`
	Phone             *string         `json:"phone"`
	Skill             *string         `json:"skill"`
	BillRate          float64         `json:"bill_rate"`
	PayoutRate        float64         `json:"payout_rate"`
	ActiveStatus      bool            `json:"active_status"`
	Notes             *string         `json:"notes"`
	AdditionalDetails json.RawMessage `json:"additional_details"`
	CreatedAt         time.Time       `json:"created_at"`
	UpdatedAt         time.Time       `json:"updated_at"`
}
