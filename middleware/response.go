package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response interface {
	Ok(value string)
	Error(status int, value string)
	Created()
	NoContent()
	Status(code int) Entity
	Header(headers Params) Header
}

type Header interface {
	Status(code int) Entity
}

type Entity interface {
	Entity(value interface{})
}

type responseBuild struct {
	ctx *gin.Context
}

type entityBuild struct {
	status int
	build  *responseBuild
}

type headerBuild struct {
	build *responseBuild
}

type Params map[string]string

func NewResponse(c *gin.Context) Response {
	return &responseBuild{ctx: c}
}

func (r *responseBuild) Error(status int, value string) {
	r.ctx.Data(status, "application/json", []byte(value))
}

func (r *responseBuild) Ok(value string) {
	r.ctx.Data(http.StatusOK, "application/json", []byte(value))
}

func (r *responseBuild) Created() {
	r.ctx.JSON(http.StatusCreated, gin.H{})
}

func (r *responseBuild) NoContent() {
	r.ctx.Writer.WriteHeader(http.StatusNoContent)
}

func (r *responseBuild) Status(code int) Entity {
	return &entityBuild{status: code, build: r}
}

func (r *responseBuild) Header(headers Params) Header {
	for key, value := range headers {
		r.ctx.Writer.Header().Add(key, value)
	}
	return &headerBuild{build: r}
}

func (e *headerBuild) Status(code int) Entity {
	return &entityBuild{status: code, build: e.build}
}

func (e *entityBuild) Entity(value interface{}) {
	e.build.ctx.JSON(e.status, value)
}
