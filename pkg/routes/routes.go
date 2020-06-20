package routes

import (
	"go-boilerplate/pkg/config"
	"go-boilerplate/pkg/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

// NewRoute NewRoute
func NewRoute(r IRouter) http.Handler {
	if config.App.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	e := gin.New()

	// middleware
	middleware.RouteMiddleware(e)

	// api
	api := e.Group("/api")
	{
		r.RegisterAPI(api)
	}

	// web
	web := e.Group("/")
	{
		r.RegisterWeb(web)
	}

	e.NoRoute(NotFound)
	e.NoMethod(NotFound)

	return e
}

// NotFound represents the 404 page.
func NotFound(ctx *gin.Context) {
	ctx.AbortWithStatusJSON(http.StatusNotFound, "PAGE NOT FOUND")
	return
}
