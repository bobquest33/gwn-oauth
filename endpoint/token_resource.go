package endpoint

import (
	"net/http"

	"github.com/gin-gonic/gin"
	. "github.com/helderfarias/gwn-oauth/log"
	"github.com/helderfarias/gwn-oauth/middleware"
)

type TokenResource struct {
	contextFactory middleware.ContextWrapperFactory
}

func (r *TokenResource) register(router *gin.Engine) {
	grupo := router.Group("/oauth")

	grupo.POST("/token", r.createToken)
}

func (r *TokenResource) createToken(c *gin.Context) {
	context := r.contextFactory.Create(c)

	service := context.GetServiceFactory().GetTokenService()

	token, err := service.Create(c.Request)
	if err != nil {
		Logger.Error("Erro: %s", err)
		context.Response().Error(http.StatusUnauthorized, err.Error())
		return
	}

	context.Response().Ok(token)
}
