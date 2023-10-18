package utils

import (
	"os"

	"github.com/rs/zerolog"
)

// InitLogger creates a zerolog logger with standardised setup to be used throughout the application
func InitLogger() *zerolog.Logger {
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	return &logger
}
