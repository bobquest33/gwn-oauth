package endpoint

import (
	"github.com/gin-gonic/gin"
	"github.com/helderfarias/gwn-oauth/middleware"
)

type RootResource struct {
	contextFactory middleware.ContextWrapperFactory
}

func (r *RootResource) register(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		context := r.contextFactory.Create(c)

		context.Response().Ok("GWN-OAuthServer")
	})
}
