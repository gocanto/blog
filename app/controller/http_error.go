package controller

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
)

type HttpError struct {
	Code             int
	Message          string
	Err              error
	ValidationErrors map[string]any
}

func MakeHttpError(code int, message string, err error) *HttpError {
	return &HttpError{
		Code:             code,
		Message:          message,
		Err:              err,
		ValidationErrors: make(map[string]any),
	}
}

func BadRequest(message string, err error) *HttpError {
	return MakeHttpError(http.StatusBadRequest, message, err)
}

func InternalServerError(message string, err error) *HttpError {
	return MakeHttpError(http.StatusInternalServerError, message, err)
}

func RespondWithErrors(message string, validationErrors map[string]any, err error) *HttpError {
	return &HttpError{
		Code:             http.StatusForbidden,
		Message:          message,
		Err:              err,
		ValidationErrors: validationErrors,
	}
}

func Unauthorised(message string, err error) *HttpError {
	return &HttpError{
		Code:             http.StatusUnauthorized,
		Message:          message,
		Err:              err,
		ValidationErrors: make(map[string]any),
	}
}

func (e *HttpError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}

	return e.Message
}

func (e *HttpError) Unwrap() error {
	return e.Err
}

func (e *HttpError) Respond(w http.ResponseWriter) {
	slog.Error("HTTP Error", "status", e.Code, "message", e.Message, "error", e.Err, "validation_errors", e.ValidationErrors)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff") // Basic security header
	w.WriteHeader(e.Code)

	payload := map[string]any{
		"message": e.Message,
	}

	if len(e.ValidationErrors) > 0 {
		payload["errors"] = e.ValidationErrors
	}

	if err := json.NewEncoder(w).Encode(payload); err != nil {
		slog.Error("Error encoding error response", "encode_error", err, "original_error", e)
		_, _ = fmt.Fprintf(w, `{"message":"Error generating error response"}`)
	}
}
