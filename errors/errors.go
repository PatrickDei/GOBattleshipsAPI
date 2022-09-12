package errors

import "net/http"

const ErrorBodyErrorCode = "error-code"
const ErrorBodyErrorArg = "error-arg"

const ErrorCodePrefix = "error."

type ErrorBody map[string]string

func NewErrorBody(code string, arg string) ErrorBody {
	return map[string]string{
		ErrorBodyErrorCode: code,
		ErrorBodyErrorArg:  arg,
	}
}

type AppError struct {
	Code int
	Body ErrorBody
}

func (ae AppError) AsResponseMessage() map[string]string {
	return ae.Body
}

func NewConflictError(body ErrorBody) *AppError {
	return &AppError{
		Code: http.StatusConflict,
		Body: body,
	}
}

func NewNotFoundError(body ErrorBody) *AppError {
	return &AppError{
		Code: http.StatusNotFound,
		Body: body,
	}
}

func NewInternalServerError(body ErrorBody) *AppError {
	return &AppError{
		Code: http.StatusInternalServerError,
		Body: body,
	}
}
