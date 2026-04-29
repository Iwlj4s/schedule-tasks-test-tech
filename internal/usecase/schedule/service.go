package schedule

import (
	"context"
	"fmt"
	"time"

	exceptionsdomain "example.com/taskservice/internal/domain/exceptions"
	scheduledomain "example.com/taskservice/internal/domain/schedule"
)

type Service struct {
	repo Repository
	now  func() time.Time
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
		now:  func() time.Time { return time.Now().UTC() },
	}
}

func (s *Service) CreateSchedule(ctx context.Context, input CreateScheduleInput) (*scheduledomain.Schedule, error) {
	normalized, err := validateCreateScheduleInput(input)
	if err != nil {
		return nil, err
	}

	model := &scheduledomain.Schedule{
		Title:            normalized.Title,
		Description:      normalized.Description,
		Status:           normalized.Status,
		RecurrenceType:   normalized.RecurrenceType,
		RecurrenceConfig: normalized.RecurrenceConfig,
		NextRunAt:        normalized.NextRunAt,
	}

	now := s.now()
	model.CreatedAt = now
	model.UpdatedAt = now

	created, err := s.repo.CreateSchedule(ctx, model)
	if err != nil {
		return nil, err
	}

	return created, nil
}

func (s *Service) DeleteSchedule(ctx context.Context, id int64) error {
	if id <= 0 {
		return fmt.Errorf("%w: id must be positive", exceptionsdomain.ErrInvalidInput) // Знаю, что пока этой функции нет
	}

	return s.repo.DeleteSchedule(ctx, id)
}

func (s *Service) GetScheduleByID(ctx context.Context, id int64) (*scheduledomain.Schedule, error) {
	if id <= 0 {
		return nil, fmt.Errorf("%w: id must be positive", exceptionsdomain.ErrInvalidInput)
	}

	return s.repo.GetScheduleByID(ctx, id)
}
