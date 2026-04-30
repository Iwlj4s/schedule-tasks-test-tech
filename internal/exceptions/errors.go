package exceptions

import (
	"errors"
	"net/http"

	helper "example.com/taskservice/internal/helpers"
)

var (
	ErrNotFound     = errors.New("task not found")
	ErrInvalidInput = errors.New("invalid input")
)

func WriteError(w http.ResponseWriter, status int, err error) {
	helper.WriteJSON(w, status, map[string]string{
		"error": err.Error(),
	})
}

func WriteUsecaseError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, ErrNotFound):
		WriteError(w, http.StatusNotFound, err)
	case errors.Is(err, ErrInvalidInput):
		WriteError(w, http.StatusBadRequest, err)
	default:
		WriteError(w, http.StatusInternalServerError, err)
	}
}
