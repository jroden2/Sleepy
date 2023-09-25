package private

import (
	"github.com/gin-gonic/gin"
)

func Routes(route *gin.RouterGroup) {
	exampleRouteGroup := route.Group("")
	{
		controller := NewExampleController(nil)
		exampleRouteGroup.GET("/", controller.ExposedFunction)
	}
}