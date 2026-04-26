package schedule

import (
	"encoding/json"
	"time"
)

type Status string
type RecurrenceType string
type RecurrenceConfig json.RawMessage

const (
	StatusNew        Status = "new"
	StatusInProgress Status = "in_progress"
	StatusDone       Status = "done"
)

const (
	RecurrenceDaily    RecurrenceType = "daily"
	RecurrenceMounthly RecurrenceType = "mounthly"
	RecurrenceSpecific RecurrenceType = "specific_dates"
	RecurrenceOddEven  RecurrenceType = "odd_even"
)

type Schedule struct {
	ID               int64            `json:"id"`
	Title            string           `json:"title"`
	Description      string           `json:"description"`
	Status           Status           `json:"status"`
	RecurrenceType   RecurrenceType   `json:"recurrence_type"`
	RecurrenceConfig RecurrenceConfig `json:"recurrence_config"`
	NextRunAt        time.Time        `json:"next_run_at"`
	CreatedAt        time.Time        `json:"created_at"`
	UpdatedAt        time.Time        `json:"updated_at"`
}

func (s Status) Valid() bool {
	switch s {
	case StatusNew, StatusInProgress, StatusDone:
		return true
	default:
		return false
	}
}

func (rt RecurrenceType) Valid() bool {
	switch rt {
	case RecurrenceDaily, RecurrenceMounthly, RecurrenceSpecific, RecurrenceOddEven:
		return true
	default:
		return false
	}
}
