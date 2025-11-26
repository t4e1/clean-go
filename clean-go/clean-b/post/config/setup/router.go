package setup

import (
	"github.com/gin-gonic/gin"
	"github.com/t4e1/clean-go/clean-b/post/adapter/in"
)

func SetupRouter(port *in.RESTInAdapter) *gin.Engine {
	r := gin.Default()

	// REST Api routing
	r.GET("/post/:id", port.GetPost)
	r.GET("/post", port.GetPosts)
	r.POST("/post", port.NewPost)
	// PUT: entire update	- can be set default value if it dosen't have enough value
	// PATCH: partial update - remain old values except input value
	r.PATCH("/post/:id", port.UpdatePost)
	r.DELETE("/post/:id", port.DeletePost)

	return r
}
