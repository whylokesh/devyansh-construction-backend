package models

import (
	"time"
)

type AttendanceStatus string

const (
	AttendancePresent AttendanceStatus = "present"
	AttendanceAbsent  AttendanceStatus = "absent"
	AttendanceHalfDay AttendanceStatus = "half_day"
)

type Attendance struct {
	ID        int              `json:"id"`
	WorkerID  int              `json:"worker_id"`
	SiteID    int              `json:"site_id"`
	Date      time.Time        `json:"date"` // Using time.Time for DATE, usually truncated to midnight
	Status    AttendanceStatus `json:"status"`
	Note      *string          `json:"note"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`
}
