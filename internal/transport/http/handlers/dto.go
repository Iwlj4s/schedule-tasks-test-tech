package handlers

import (
	"encoding/json"
	"time"

	scheduledomain "example.com/taskservice/internal/domain/schedule"
	taskdomain "example.com/taskservice/internal/domain/task"
)

// Tasks DTO
type taskMutationDTO struct {
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Status      taskdomain.Status `json:"status"`
}

type taskDTO struct {
	ID          int64             `json:"id"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Status      taskdomain.Status `json:"status"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
}

func newTaskDTO(task *taskdomain.Task) taskDTO {
	return taskDTO{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	}
}

// Schedule DTO
type scheduleMutationDTO struct {
	Title            string          `json:"title"`
	Description      string          `json:"description"`
	Status           string          `json:"status"`
	RecurrenceType   string          `json:"recurrence_type"`
	RecurrenceConfig json.RawMessage `json:"recurrence_config"`
	NextRunAt        time.Time       `json:"next_run_at"`
}

type scheduleDTO struct {
	ID               int64           `json:"id"`
	Title            string          `json:"title"`
	Description      string          `json:"description"`
	Status           string          `json:"status"`
	RecurrenceType   string          `json:"recurrence_type"`
	RecurrenceConfig json.RawMessage `json:"recurrence_config"`
	NextRunAt        time.Time       `json:"next_run_at"`
	CreatedAt        time.Time       `json:"created_at"`
	UpdatedAt        time.Time       `json:"updated_at"`
}

func newScheduleDTO(schedule *scheduledomain.Schedule) scheduleDTO {
	return scheduleDTO{
		ID:               schedule.ID,
		Title:            schedule.Title,
		Description:      schedule.Description,
		Status:           string(schedule.Status),
		RecurrenceType:   string(schedule.RecurrenceType),
		RecurrenceConfig: json.RawMessage(schedule.RecurrenceConfig),
		NextRunAt:        schedule.NextRunAt,
		CreatedAt:        schedule.CreatedAt,
		UpdatedAt:        schedule.UpdatedAt,
	}
}
