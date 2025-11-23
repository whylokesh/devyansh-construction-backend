package models

import (
	"encoding/json"
	"time"
)

type SiteSummarySnapshot struct {
	ID          int             `json:"id"`
	SiteID      int             `json:"site_id"`
	SummaryJSON json.RawMessage `json:"summary_json"`
	CreatedAt   time.Time       `json:"created_at"`
}
