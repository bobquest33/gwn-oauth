package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/helderfarias/gwn-oauth/service"
	"github.com/helderfarias/gwn-oauth/util"
	"github.com/jmoiron/sqlx"
)

type ContextWrapperFactory interface {
	Create(context *gin.Context) ContextWrapper
}

type ContextWrapper interface {
	Response() Response
	GetParam(name string) string
	GetParamAsInt(name string) int
	GetServiceFactory() service.ServiceFactory
}

type contextWrapper struct {
	parserForm     bool
	context        *gin.Context
	response       Response
	serviceFactory service.ServiceFactory
}

type contextWrapperFactory struct {
}

func NewContextWrapperFactory() ContextWrapperFactory {
	return &contextWrapperFactory{}
}

func (ctf *contextWrapperFactory) Create(context *gin.Context) ContextWrapper {
	wrapper := &contextWrapper{}
	wrapper.context = context
	wrapper.response = NewResponse(context)
	wrapper.serviceFactory = ctf.createServiceFactory(context)
	return wrapper
}

func (c *contextWrapper) Response() Response {
	return c.response
}

func (c *contextWrapper) GetParam(name string) string {
	c.prepareParams()

	return c.context.Request.Form.Get(name)
}

func (c *contextWrapper) GetParamAsInt(name string) int {
	value := c.GetParam(name)

	if value == "" {
		return 0
	}

	return util.ToInteger(value)
}

func (c *contextWrapper) GetServiceFactory() service.ServiceFactory {
	return c.serviceFactory
}

func (c *contextWrapper) prepareParams() {
	if !c.parserForm {
		c.context.Request.ParseForm()
		c.parserForm = true
	}
}

func (c *contextWrapperFactory) createServiceFactory(context *gin.Context) service.ServiceFactory {
	return service.NewServiceFactory(context.MustGet("db").(*sqlx.DB))
}
