package controllers

import (
	"github.com/gofiber/fiber"
	"github.com/jroden2/Sleepy/Fiber/pkg/controllers/private"
	"github.com/jroden2/Sleepy/Fiber/utils"
	"github.com/rs/zerolog"
)

func Initialise(logger *zerolog.Logger) {
	if logger == nil {
		logger = utils.InitLogger()
	}

	router := fiber.New()

	router.Group("")
	{
		private.Routes(router, logger)
	}

	router.Listen(":8080")
}
