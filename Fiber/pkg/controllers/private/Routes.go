package private

import "github.com/gofiber/fiber"

func Routes(route fiber.Router) {

	exampleRouteGroup := route.Group("base")
	{

		exampleRouteGroup.Get("/greet")
	}

}