package routes

import (
	"go-boilerplate/internal/pkg/middleware"
	"go-boilerplate/pkg/config"
	"net/http"

	"github.com/gin-contrib/pprof"
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
	if config.App.Debug {
		pprof.Register(e)
	}

	// middleware
	middleware.RouteMiddleware(e)

	// api
	api := e.Group("/api")
	{
		r.RegisterAPI(api)
	}

	// websocket
	ws := e.Group("/ws")
	{
		r.RegisterWebSocket(ws)
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
