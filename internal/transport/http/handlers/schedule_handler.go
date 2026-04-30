package handlers

import (
	"net/http"

	scheduledomain "example.com/taskservice/internal/domain/schedule"
	exceptions "example.com/taskservice/internal/exceptions"
	helpers "example.com/taskservice/internal/helpers"
	scheduleusecase "example.com/taskservice/internal/usecase/schedule"
)

type ScheduleHandler struct {
	usecase scheduleusecase.ScheduleUseCase
}

func NewScheduleHandler(usecase scheduleusecase.ScheduleUseCase) *ScheduleHandler {
	return &ScheduleHandler{usecase: usecase}
}

func (h *ScheduleHandler) CreateSchedule(w http.ResponseWriter, r *http.Request) {
	var req scheduleMutationDTO
	if err := helpers.DecodeJSON(r, &req); err != nil {
		exceptions.WriteError(w, http.StatusBadRequest, err)
		return
	}

	created, err := h.usecase.CreateSchedule(r.Context(), scheduleusecase.CreateScheduleInput{
		Title:            req.Title,
		Description:      req.Description,
		Status:           scheduledomain.Status(req.Status),
		RecurrenceType:   scheduledomain.RecurrenceType(req.RecurrenceType),
		RecurrenceConfig: scheduledomain.RecurrenceConfig(req.RecurrenceConfig),
		NextRunAt:        req.NextRunAt,
	})

	if err != nil {
		exceptions.WriteUsecaseError(w, err)
		return
	}

	helpers.WriteJSON(w, http.StatusCreated, newScheduleDTO(created))
}

func (h *ScheduleHandler) DeleteSchedule(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromRequest(r)
	if err != nil {
		exceptions.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := h.usecase.DeleteSchedule(r.Context(), id); err != nil {
		exceptions.WriteUsecaseError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *ScheduleHandler) GetScheduleByID(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromRequest(r)
	if err != nil {
		exceptions.WriteError(w, http.StatusBadRequest, err)
		return
	}
	schedule, err := h.usecase.GetScheduleByID(r.Context(), id)
	if err != nil {
		exceptions.WriteUsecaseError(w, err)
		return
	}

	helpers.WriteJSON(w, http.StatusOK, newScheduleDTO(schedule))
}
