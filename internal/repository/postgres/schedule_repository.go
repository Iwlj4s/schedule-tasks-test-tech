package postgres

import (
	"context"
	"errors"

	scheduledomain "example.com/taskservice/internal/domain/schedule"
)

func (r *Repository) CreateSchedule(ctx context.Context, s *scheduledomain.Schedule) (*scheduledomain.Schedule, error) {
	return nil, errors.New("not implemented")
}

func (r *Repository) GetScheduleByID(ctx context.Context, id int64) (*scheduledomain.Schedule, error) {
	const query = `
		SELECT id, title, description, status, recurrence_type, recurrence_config, next_run_at, created_at, updated_at
		FROM task_schedules
		WHERE id = $1
	`
	return nil, errors.New("not implemented")
}
