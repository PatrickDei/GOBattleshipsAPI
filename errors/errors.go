package errors

import "net/http"

type AppError struct {
	Code int
	Body map[string]string
}

func (ae AppError) AsResponseMessage() map[string]string {
	return ae.Body
}

func NewConflictError(body map[string]string) *AppError {
	return &AppError{
		Code: http.StatusConflict,
		Body: body,
	}
}

func NewNotFoundError(body map[string]string) *AppError {
	return &AppError{
		Code: http.StatusNotFound,
		Body: body,
	}
}

func NewInternalServerError(body map[string]string) *AppError {
	return &AppError{
		Code: http.StatusInternalServerError,
		Body: body,
	}
}
