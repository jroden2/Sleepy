package private

import (
	"github.com/gofiber/fiber"
	"github.com/rs/zerolog"
)

func Routes(route fiber.Router, logger *zerolog.Logger) {

	exampleRouteGroup := route.Group("")
	{
		controller := NewBaseController(logger)
		exampleRouteGroup.Get("/", controller.Greet)
	}

}
