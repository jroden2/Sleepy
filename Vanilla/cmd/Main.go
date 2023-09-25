package main

import (
	"github.com/joho/godotenv"
	"github.com/jroden2/Sleepy/Vanilla/pkg/controllers"
	"github.com/jroden2/Sleepy/Vanilla/utils"
)

func main() {
	logger := utils.InitLogger()

	err := godotenv.Load()
	if err != nil {
		logger.Error().Msg("Error loading .env file")
		return
	}

	controllers.NewServer(logger)
}