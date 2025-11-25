package setup

import (
	"github.com/gin-gonic/gin"
	"github.com/t4e1/clean-go/clean-b/post/adapter/in"
)

func SetupRouter(port *in.RESTInAdapter) *gin.Engine {
	r := gin.Default()

	// REST Api routing
	r.GET("/post/:id", port.GetPost)

	return r
}
