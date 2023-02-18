package handlers

import (
	"errors"
	"net/http"

	"github.com/kotaroyamazaki/go-clean-arch-sample-with-standard/internal/entities"
)

func GetStatusCode(err error) int {
	switch {
	case errors.Is(err, entities.ErrNotFound), errors.Is(err, entities.ErrDuplicatedPrimaryKey):
		return http.StatusInternalServerError
	case errors.Is(err, entities.ErrValidattion):
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
