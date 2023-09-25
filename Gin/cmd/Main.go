package main

import (
	"github.com/joho/godotenv"
	"github.com/jroden2/Sleepy/Gin/pkg/controllers"
	"github.com/jroden2/Sleepy/Gin/utils"
)

func main() {
	logger := utils.InitLogger()

	err := godotenv.Load()
	if err != nil {
		logger.Error().Msg("Error loading .env file")
		return
	}
	controllers.Initialise(logger)
}