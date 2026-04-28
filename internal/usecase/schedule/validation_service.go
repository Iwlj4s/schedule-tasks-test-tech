package schedule

import (
	"fmt"
	"strings"

	exceptionsdomain "example.com/taskservice/internal/domain/exceptions"
	scheduledomain "example.com/taskservice/internal/domain/schedule"
)

func validateCreateScheduleInput(input CreateScheduleInput) (CreateScheduleInput, error) {
	input.Title = strings.TrimSpace(input.Title)
	input.Description = strings.TrimSpace(input.Description)

	if input.Title == "" {
		return CreateScheduleInput{}, fmt.Errorf("%w: title is required", exceptionsdomain.ErrInvalidInput)
	}

	if input.Status == "" {
		input.Status = scheduledomain.StatusNew
	}

	if !input.Status.Valid() {
		return CreateScheduleInput{}, fmt.Errorf("%w: invalid status", exceptionsdomain.ErrInvalidInput)
	}

	if !input.RecurrenceType.Valid() {
		return CreateScheduleInput{}, fmt.Errorf("%w: invalid recurrence type", exceptionsdomain.ErrInvalidInput)
	}

	if input.NextRunAt.IsZero() {
		return CreateScheduleInput{}, fmt.Errorf("%w: next_run_at is required", exceptionsdomain.ErrInvalidInput)
	}

	return input, nil
}
