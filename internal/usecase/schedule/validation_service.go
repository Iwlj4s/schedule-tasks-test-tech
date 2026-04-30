package schedule

import (
	"fmt"
	"strings"

	scheduledomain "example.com/taskservice/internal/domain/schedule"
	exceptions "example.com/taskservice/internal/exceptions"
)

func validateCreateScheduleInput(input CreateScheduleInput) (CreateScheduleInput, error) {
	input.Title = strings.TrimSpace(input.Title)
	input.Description = strings.TrimSpace(input.Description)

	if input.Title == "" {
		return CreateScheduleInput{}, fmt.Errorf("%w: title is required", exceptions.ErrInvalidInput)
	}

	if input.Status == "" {
		input.Status = scheduledomain.StatusNew
	}

	if !input.Status.Valid() {
		return CreateScheduleInput{}, fmt.Errorf("%w: invalid status", exceptions.ErrInvalidInput)
	}

	if !input.RecurrenceType.Valid() {
		return CreateScheduleInput{}, fmt.Errorf("%w: invalid recurrence type", exceptions.ErrInvalidInput)
	}

	if input.NextRunAt.IsZero() {
		return CreateScheduleInput{}, fmt.Errorf("%w: next_run_at is required", exceptions.ErrInvalidInput)
	}

	return input, nil
}
