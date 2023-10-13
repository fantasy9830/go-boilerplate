package http

import (
	"go-boilerplate/pkg/config"
	"go-boilerplate/pkg/middleware"
	"go-boilerplate/pkg/net/http/server"
	"go-boilerplate/pkg/version"
	"net/http"

	"github.com/gin-gonic/gin"
)

var _ Router = (*Route)(nil)

type Router interface {
	Handler() http.Handler
}

// Route
type Route struct {
	UserHandler *UserHandler
}

func NewRouter(userHandler *UserHandler) Router {
	return &Route{
		UserHandler: userHandler,
	}
}

func Init() {
	router := InitializeRouter()
	srv := server.NewServer(router.Handler())

	srv.Start()
}

// Handler load initializes the routing of the application.
func (r *Route) Handler() http.Handler {
	if config.App.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	e := gin.New()

	e.Use(middleware.Cors())
	e.Use(gin.Recovery())
	e.Use(gin.Logger())

	api := e.Group("/api")
	{
		api.GET("/version", version.APIVersion)
		api.GET("/healthz", func(ctx *gin.Context) {
			ctx.AbortWithStatus(http.StatusOK)
			ctx.String(http.StatusOK, "ok")
		})
	}

	return e
}
