package httperrors

import (
	"net/http"
	"strings"
)

const (
	ErrInvalidJSON   = "Missed requried param or it is null"
	ErrNotPositiveId = "Id must be greater than 0"
	ErrNotFound      = "Object not found in database"
	ErrIdNotInteger  = "Id must be integer"
	ErrInternal      = "Internal error, please contact support"
)

type ErrorResponse struct {
	Error  string `json:"error"`
	Detail string `json:"detail"`
}

func NewErrorResponse(errorMessage string) (int, *ErrorResponse) {
	statusCode, errorDescription := parseError(errorMessage)
	return statusCode, &ErrorResponse{Error: errorDescription, Detail: errorMessage}
}

func parseError(errorMessage string) (int, string) {
	switch {
	case strings.Contains(errorMessage, "Field validation"):
		return http.StatusBadRequest, ErrInvalidJSON
	case strings.Contains(errorMessage, "Not positive id"):
		return http.StatusBadRequest, ErrNotPositiveId
	case strings.Contains(errorMessage, "ParseInt"):
		return http.StatusBadRequest, ErrIdNotInteger
	case strings.Contains(errorMessage, "no rows in result set"):
		return http.StatusNotFound, ErrNotFound
	default:
		return http.StatusInternalServerError, ErrInternal
	}
}
