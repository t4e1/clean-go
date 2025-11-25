package in

import (
	"github.com/gin-gonic/gin"
	"github.com/t4e1/clean-go/clean-b/post/internal/ports/in"
)

// packages for REST API implementation

type RESTInAdapter struct {
	inboundPort in.InboundRESTApiPort
}

func InitRESTAdapter(inPort in.InboundRESTApiPort) *RESTInAdapter {
	return &RESTInAdapter{
		inboundPort: inPort,
	}
}

func (r *RESTInAdapter) GetPost(c *gin.Context) {
	c.JSON(200, "ok")
}

func (r *RESTInAdapter) GetPosts(c *gin.Context) {
	c.JSON(200, "ok")
}

func (r *RESTInAdapter) NewPost(c *gin.Context) {
	c.JSON(200, "ok")
}

func (r *RESTInAdapter) UpdatePost(c *gin.Context) {
	c.JSON(200, "ok")
}

func (r *RESTInAdapter) DeletePost(c *gin.Context) {
	c.JSON(200, "ok")
}
