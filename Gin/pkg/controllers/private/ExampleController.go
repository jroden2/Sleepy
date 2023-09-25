package private

import (
	"github.com/gin-gonic/gin"
	"github.com/jroden2/Sleepy/Gin/pkg/services"
)

type exampleController struct {
	es services.ExampleService
}

func NewExampleController(es *services.ExampleService) ExampleController {
	if es == nil {
		tmp := services.NewExampleService()
		es = &tmp
	}

	return &exampleController{
		es: *es,
	}
}

type ExampleController interface {
	ExposedFunction(ctx *gin.Context)
}

// ExposedFunction is our endpoint handler, accepts a gin context
func (c *exampleController) ExposedFunction(ctx *gin.Context) {
	// we can now call the exposed function and return it to the context.
	ctx.JSON(c.es.ExposedFunction())
}