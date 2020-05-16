package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"

	"github.com/gin-gonic/gin"
	"github.com/codechaitu/api/config"
	"github.com/codechaitu/api/errors"
	"go.uber.org/zap"
)

type ServerImpl interface {
	CreateRouter() *gin.Engine
	Start(*gin.Engine) errors.ExitCode
}

type server struct {
	logger *zap.Logger
	env    *config.Env
}

func New(logger *zap.Logger, env *config.Env) ServerImpl {
	return &server{
		logger: logger,
		env:    env,
	}
}

// CreateRouter creates and configures a server
func (s *server) CreateRouter() *gin.Engine {
	router := gin.Default()
	setupRoutes(router)
	return router
}

// Start starts servers and handles shutdown gracefully
func (s *server) Start(router *gin.Engine) errors.ExitCode {
	port := s.env.Port
	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	// Listening to connections
	go func() {
		if err := server.ListenAndServe(); err != nil {
			s.logger.Error("error when server tyring to listen and serve on given port:" + port)
		}
	}()

	// Graceful shutdown
	exit := make(chan os.Signal)
	signal.Notify(exit, os.Interrupt)
	<-exit
	s.logger.Info("input signal received to terminate")

	// Allowing time for graceful degradation
	ctx, cancel := context.WithTimeout(context.Background(), s.env.ServerDegradationTime)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		s.logger.Error("error while shutting down the server, " + zap.Error(err).String)
		return errors.Error
	}

	s.logger.Info("server is now off")
	return errors.Nil
}
