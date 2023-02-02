package errors

import "net/http"

type AppError struct {
	Code    int
	Message string
}

func NewNotFoundError(m string) *AppError {
	return &AppError{Code: http.StatusNotFound, Message: m}
}

func NewUnExpectedError(m string) *AppError {
	return &AppError{Code: http.StatusInternalServerError, Message: m}
}

func NewValidationError(m string) *AppError {
	return &AppError{Code: http.StatusUnprocessableEntity, Message: m}
}
