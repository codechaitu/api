package main

import (
	"os"

	"github.com/codechaitu/api/config"
	"github.com/codechaitu/api/errors"
	"github.com/codechaitu/api/server"
	"go.uber.org/zap"
)

func main() {
	// Start a logger for logging
	logger, err := zap.NewProduction()
	if err != nil {
		os.Exit(int(errors.LoggerError))
	}
	env, err := config.ReadEnvVariables()
	if err != nil {
		logger.Error("error to read environment variables: " + zap.Error(err).String)
		os.Exit(int(errors.Error))
	}

	server := server.New(logger, env)
	exitCode := server.Start(server.CreateRouter())
	os.Exit(int(exitCode))
}
