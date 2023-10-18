package private

import (
	"github.com/gofiber/fiber"
	"github.com/jroden2/Sleepy/Fiber/pkg/services"
	"github.com/jroden2/Sleepy/Fiber/utils"
	"github.com/rs/zerolog"
)

type baseController struct {
	logger zerolog.Logger
	es     services.ExampleService
}

func NewBaseController(logger *zerolog.Logger) BaseController {
	if logger == nil {
		logger = utils.InitLogger()
	}

	return &baseController{logger: *logger, es: services.NewExampleService()}
}

type BaseController interface {
}

func (c *baseController) Greet(ctx *fiber.Ctx) {
	status, body := c.es.ExposedFunction()
	ctx.Status(status)
	ctx.JSON(body)
}
