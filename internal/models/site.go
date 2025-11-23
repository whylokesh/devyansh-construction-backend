package models

import (
	"encoding/json"
	"time"
)

type SiteStatus string

const (
	SiteStatusActive    SiteStatus = "active"
	SiteStatusCompleted SiteStatus = "completed"
)

type Site struct {
	ID                int             `json:"id"`
	Name              string          `json:"name"`
	ClientName        string          `json:"client_name"`
	ClientPhone       *string         `json:"client_phone"`
	Location          *string         `json:"location"`
	SiteDocuments     *string         `json:"site_documents"`
	AdditionalDetails json.RawMessage `json:"additional_details"`
	StartDate         time.Time       `json:"start_date"`
	EndDate           *time.Time      `json:"end_date"`
	Status            SiteStatus      `json:"status"`
	CreatedAt         time.Time       `json:"created_at"`
	UpdatedAt         time.Time       `json:"updated_at"`
}
