package app

import (
	"errors"
	"fmt"
	"net/http"
)

type AppError struct {
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e *AppError) Error() string {
	return e.Message
}

func ConstructBadRequestError(code, message string) *AppError {
	if message == "" {
		message = "Bad Request"
	}

	return &AppError{
		Status:  http.StatusBadRequest,
		Code:    code,
		Message: message,
	}
}

func ConstructForbiddenError(code, message string) *AppError {
	if message == "" {
		message = "Forbidden"
	}

	return &AppError{
		Status:  http.StatusForbidden,
		Code:    code,
		Message: message,
	}
}

func ConstructNotFoundError(code, message string) *AppError {
	if message == "" {
		message = "Not Found"
	}

	return &AppError{
		Status:  http.StatusNotFound,
		Code:    code,
		Message: message,
	}
}

func ConstructConflictError(code, message string) *AppError {
	if message == "" {
		message = "Conflict"
	}

	return &AppError{
		Status:  http.StatusConflict,
		Code:    code,
		Message: message,
	}
}

func ConstructUnauthorizedError(code, message string) *AppError {
	if message == "" {
		message = "Unauthorized"
	}

	return &AppError{
		Status:  http.StatusUnauthorized,
		Code:    code,
		Message: message,
	}
}

func ConstructPaymentRequiredError(code, message string) *AppError {
	if message == "" {
		message = "Payment Required"
	}

	return &AppError{
		Status:  http.StatusPaymentRequired,
		Code:    code,
		Message: message,
	}
}

func ConstructUnprocessableEntityError(code, message string) *AppError {
	if message == "" {
		message = "Unprocessable Entity"
	}

	return &AppError{
		Status:  http.StatusUnprocessableEntity,
		Code:    code,
		Message: message,
	}
}

func ConstructInternalServerError(code, message string) *AppError {
	if message == "" {
		message = "Internal Server Error"
	}

	return &AppError{
		Status:  http.StatusInternalServerError,
		Code:    code,
		Message: message,
	}
}

func WrapError(code, message string, err error) *AppError {
	if message == "" {
		message = "Error occurred"
	}

	if err == nil {
		err = errors.New("internal server error")
	}

	return &AppError{
		Status:  http.StatusInternalServerError,
		Code:    code,
		Message: fmt.Sprintf("%s: %s", message, err.Error()),
	}
}

func CustomError(status int, code, message string) *AppError {
	return &AppError{
		Status:  status,
		Code:    code,
		Message: message,
	}
}
