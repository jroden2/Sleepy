package controllers

import (
	"fmt"
	"github.com/rs/zerolog"
	"net/http"
	"os"
)

func NewServer(logger *zerolog.Logger) {
	// Set the port
	PORT := os.Getenv("ServerPort")
	if PORT == "" {
		PORT = "8080"
	}
	hostPort := fmt.Sprintf(":%s", PORT)
	logger.Info().Msg(hostPort)

	h := NewHandler(nil)
	http.HandleFunc("/", h.ExampleHandler)

	err := http.ListenAndServe(hostPort, nil)
	if err != nil {
		logger.Error().Msgf("something went wrong: %v", err.Error())
	}
}