package postgres

import (
	"context"
	"errors"

	"example.com/taskservice/internal/domain/exceptions"
	scheduledomain "example.com/taskservice/internal/domain/schedule"
	"github.com/jackc/pgx/v5"
)

func (r *Repository) CreateSchedule(ctx context.Context, s *scheduledomain.Schedule) (*scheduledomain.Schedule, error) {
	const query = `
		INSERT INTO task_schedules (title, description, status, recurrence_type, recurrence_config, next_run_at, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING title, description, status, recurrence_type, recurrence_config, next_run_at, created_at, updated_at
	`

	row := r.pool.QueryRow(
		ctx,
		query,
		s.Title,
		s.Description,
		s.Status,
		s.RecurrenceType,
		s.RecurrenceConfig,
		s.NextRunAt,
		s.CreatedAt,
		s.UpdatedAt)

	created, err := scanSchedule(row)
	if err != nil {
		return nil, err
	}

	return created, nil
}

func (r *Repository) DeleteSchedule(ctx context.Context, id int64) error {
	const query = `DELETE FROM task_schedules WHERE id = $1`

	result, err := r.pool.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return exceptions.ErrNotFound
	}

	return nil
}

func (r *Repository) GetScheduleByID(ctx context.Context, id int64) (*scheduledomain.Schedule, error) {
	const query = `
		SELECT id, title, description, status, recurrence_type, recurrence_config, next_run_at, created_at, updated_at
		FROM task_schedules
		WHERE id = $1
	`

	row := r.pool.QueryRow(ctx, query, id)
	found, err := scanSchedule(row)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, exceptions.ErrNotFound
		}
		return nil, err
	}

	return found, nil
}
