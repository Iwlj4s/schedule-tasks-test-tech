package postgres

import (
	scheduledomain "example.com/taskservice/internal/domain/schedule"
	taskdomain "example.com/taskservice/internal/domain/task"
)

type rowScanner interface {
	Scan(dest ...any) error
}

func scanTask(scanner rowScanner) (*taskdomain.Task, error) {
	var (
		task   taskdomain.Task
		status string
	)

	if err := scanner.Scan(
		&task.ID,
		&task.Title,
		&task.Description,
		&status,
		&task.CreatedAt,
		&task.UpdatedAt,
	); err != nil {
		return nil, err
	}

	task.Status = taskdomain.Status(status)

	return &task, nil
}

func scanSchedule(scanner rowScanner) (*scheduledomain.Schedule, error) {
	var (
		s      scheduledomain.Schedule
		status string
	)
	if err := scanner.Scan(
		&s.ID,
		&s.Title,
		&s.Description,
		&s.Status,
		&s.RecurrenceType,
		&s.RecurrenceConfig,
		&s.NextRunAt,
		&s.CreatedAt,
		&s.UpdatedAt,
	); err != nil {
		return nil, err
	}

	s.Status = scheduledomain.Status(status)

	return &s, nil
}
