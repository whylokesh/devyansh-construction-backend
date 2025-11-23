package models

import (
	"encoding/json"
	"time"
)

type RecordStatus string

const (
	RecordStatusDraft   RecordStatus = "draft"
	RecordStatusSaved   RecordStatus = "saved"
	RecordStatusPaid    RecordStatus = "paid"
	RecordStatusPartial RecordStatus = "partial"
)

type Advance struct {
	ID        int       `json:"id"`
	WorkerID  int       `json:"worker_id"`
	Amount    float64   `json:"amount"`
	GivenOn   time.Time `json:"given_on"`
	Note      *string   `json:"note"`
	CreatedAt time.Time `json:"created_at"`
}

type Payout struct {
	ID          int             `json:"id"`
	PeriodStart time.Time       `json:"period_start"`
	PeriodEnd   time.Time       `json:"period_end"`
	PayoutJSON  json.RawMessage `json:"payout_json"`
	Status      RecordStatus    `json:"status"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
}

type AdvanceApplication struct {
	ID            int       `json:"id"`
	WorkerID      int       `json:"worker_id"`
	PayoutID      int       `json:"payout_id"`
	AppliedAmount float64   `json:"applied_amount"`
	CreatedAt     time.Time `json:"created_at"`
}

type Bill struct {
	ID          int             `json:"id"`
	SiteID      int             `json:"site_id"`
	PeriodStart time.Time       `json:"period_start"`
	PeriodEnd   time.Time       `json:"period_end"`
	BillJSON    json.RawMessage `json:"bill_json"`
	Status      RecordStatus    `json:"status"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
}
