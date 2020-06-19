package routes

import (
	"go-boilerplate/pkg/config"
	"go-boilerplate/pkg/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Load Load
func Load() http.Handler {
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
		registerAPI(api)
	}

	// web
	web := e.Group("/")
	{
		registerWeb(web)
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
