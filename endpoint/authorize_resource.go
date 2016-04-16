package endpoint

import (
	"net/http"

	"github.com/gin-gonic/gin"
	. "github.com/helderfarias/gwn-oauth/log"
	"github.com/helderfarias/gwn-oauth/middleware"
)

type AuthorizeResource struct {
	contextFactory middleware.ContextWrapperFactory
}

func (r *AuthorizeResource) register(router *gin.Engine) {
	grupo := router.Group("/oauth")

	grupo.GET("/authorize", r.createAuthorizationCode)
}

func (r *AuthorizeResource) createAuthorizationCode(c *gin.Context) {
	context := r.contextFactory.Create(c)

	service := context.GetServiceFactory().GetAuthzService()

	_, err := service.Create(c.Request, c.Writer)
	if err != nil {
		Logger.Error("Erro: %s", err)
		context.Response().Error(http.StatusUnauthorized, err.Error())
		return
	}
}
