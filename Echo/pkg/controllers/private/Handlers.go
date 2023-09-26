package private

import (
	"github.com/jroden2/Sleepy/Echo/pkg/services"
	"github.com/labstack/echo/v4"
)

type exampleHandlers struct {
	es services.ExampleService
}

func NewExampleHandler(es *services.ExampleService) ExampleHandler {
	if es == nil {
		tmp := services.NewExampleService()
		es = &tmp
	}
	return &exampleHandlers{
		es: *es,
	}
}

type ExampleHandler interface {
	ExposedHandler(ctx echo.Context) error
}

func (h *exampleHandlers) ExposedHandler(ctx echo.Context) error {
	return ctx.JSON(h.es.ExposedFunction())
}
