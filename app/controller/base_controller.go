package controller

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type BaseController func(w http.ResponseWriter, r *http.Request) *HttpError

func CreateHandle(callback BaseController) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if err := callback(writer, request); err != nil {
			err.Respond(writer)
			return // Stop processing after error response.
		}

		// If the callback returns nil, it means success and the handler.
		// The caller itself is responsible for writing the success response.
	}
}

func SendJSON(writer http.ResponseWriter, statusCode int, data any) *HttpError {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	writer.WriteHeader(statusCode)

	if data == nil {
		// Handle cases where no data needs to be sent (e.g., 204 No Content)
		// Although usually, 204 doesn't have a body or Content-Type.
		// This check prevents json.NewEncoder from writing "null".
		return nil
	}

	if err := json.NewEncoder(writer).Encode(data); err != nil {
		slog.Error("Error encoding success response", "error", err)
		return InternalServerError("Failed to encode response", err)
	}

	return nil // Signal success
}
