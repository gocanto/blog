package main

import (
	"github.com/gocanto/blog/packages/core"
	"github.com/gocanto/blog/packages/user"
	"log/slog"
	"net/http"
)

func main() {
	if fileLogs, err := core.MakeDefaultFileLogs(); err != nil {
		panic("error opening file: " + err.Error())
	} else {
		defer fileLogs.Close()
	}

	validator := core.MakeValidator()

	users := user.HandleUsers{
		Validator: validator,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("POST /users", users.Create)

	slog.Info("Starting server on :8080")

	if err := http.ListenAndServe("localhost:8080", mux); err != nil {
		slog.Error("Error starting server", "error", err)
		panic("Error starting server.")
	}
}
