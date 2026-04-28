package schedule

import (
	"context"
	"time"

	scheduledomain "example.com/taskservice/internal/domain/schedule"
)

type Repository interface {
	CreateSchedule(ctx context.Context, schedule *scheduledomain.Schedule) (*scheduledomain.Schedule, error)
	DeleteSchedule(ctx context.Context, id int64) error
	GetScheduleByID(ctx context.Context, id int64) (*scheduledomain.Schedule, error)
}

type CreateScheduleInput struct {
	Title            string
	Description      string
	Status           scheduledomain.Status
	RecurrenceType   scheduledomain.RecurrenceType
	RecurrenceConfig scheduledomain.RecurrenceConfig
	NextRunAt        time.Time
}
