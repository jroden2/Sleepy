package controllers

import (
	"github.com/jroden2/Sleepy/Echo/pkg/controllers/private"
	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {

	eh := private.NewExampleHandler(nil)
	e.GET("/", eh.ExposedHandler)

}
