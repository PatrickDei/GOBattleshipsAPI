package errors

import "net/http"

const ERROR_BODY_ERROR_CODE = "error-code"
const ERROR_BODY_ERROR_ARG = "error-arg"

const ERROR_CODE_PREFIX = "error."

type ErrorBody map[string]string

func NewErrorBody(code string, arg string) ErrorBody {
	return map[string]string{
		ERROR_BODY_ERROR_CODE: code,
		ERROR_BODY_ERROR_ARG:  arg,
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
