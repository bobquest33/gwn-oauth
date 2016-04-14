package endpoint

import "github.com/gin-gonic/gin"
import "github.com/helderfarias/gwn-oauth/middleware"

type Resource interface {
	register(router *gin.Engine)
}

var endpoints []Resource

func init() {
	endpoints = make([]Resource, 0)
	factory := middleware.NewContextWrapperFactory()
	endpoints = append(endpoints, &RootResource{contextFactory: factory})
	endpoints = append(endpoints, &TokenResource{contextFactory: factory})
	endpoints = append(endpoints, &AuthorizeResource{contextFactory: factory})
}

func RegisterEndpoints(router *gin.Engine) {
	for _, r := range endpoints {
		r.register(router)
	}
}
