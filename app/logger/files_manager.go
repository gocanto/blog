package logger

import (
	"fmt"
	"github.com/gocanto/blog/app/env"
	"log/slog"
	"os"
	"time"
)

type FilesManager struct {
	path   string
	file   *os.File
	logger *slog.Logger
	env    *env.Environment
}

func MakeFilesManager(env *env.Environment) (Managers, error) {
	manager := FilesManager{}
	manager.env = env

	manager.path = manager.DefaultPath()
	resource, err := os.OpenFile(manager.path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		return FilesManager{}, err
	}

	handler := slog.New(slog.NewTextHandler(resource, nil))
	slog.SetDefault(handler)

	manager.file = resource
	manager.logger = handler

	return manager, nil
}

func (manager FilesManager) DefaultPath() string {
	logsEnvironment := manager.env.Logs

	return fmt.Sprintf(
		logsEnvironment.Dir,
		time.Now().UTC().Format(logsEnvironment.DateFormat),
	)
}

func (manager FilesManager) Close() bool {
	if err := manager.file.Close(); err != nil {
		manager.logger.Error("error closing file: " + err.Error())

		return false
	}

	return true
}
