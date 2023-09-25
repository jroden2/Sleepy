package controllers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jroden2/Sleepy/Gin/pkg/controllers/private"
	"github.com/rs/zerolog"

	"net/http"
	"os"
	"os/signal"
	"time"
)

func Initialise(logger *zerolog.Logger) {
	router := gin.New()
	gin.SetMode(gin.DebugMode)

	PORT := os.Getenv("ServerPort")
	if PORT == "" {
		PORT = "8080"
	}
	hostPort := fmt.Sprintf(":%s", PORT)
	logger.Info().Msg(hostPort)

	router.Use(gin.Recovery())

	base := router.Group("")
	{
		private.Routes(base)
	}

	srv := &http.Server{
		Addr:    hostPort,
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			logger.Error().Msgf("listen: %s", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	logger.Info().Msg("Shutdown Server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Info().Msgf("Server Shutdown: %v", err)
	}

	logger.Info().Msg("Server exiting")
}