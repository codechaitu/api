package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
)

// CreateRouter creates and configures a server
func CreateRouter() *gin.Engine {
	router := gin.Default()
	setupRoutes(router)
	return router
}

// StartServer starts servers and handles shutdown gracefully
func StartServer(router *gin.Engine) {
	server := &http.Server{
		// TODO: config use
		Addr:    ":8080",
		Handler: router,
	}

	// Listening to connections
	go func() {
		if err := server.ListenAndServe(); err != nil {
			// TODO: Handle error and log
		}
	}()

	// Graceful shutdown
	exit := make(chan os.Signal)
	signal.Notify(exit, os.Interrupt)
	<-exit
	// TODO: Handle error and log

	// Allowing 10 seconds for graceful degradation
	// TODO: Time exit with config file
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		// TODO: logger with error
	}
	// TODO : log system shutdown
}
