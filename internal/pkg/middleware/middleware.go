package middleware

import (
	"go-boilerplate/pkg/config"

	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

// RouteMiddleware 設定 Middleware
func RouteMiddleware(e *gin.Engine) *gin.Engine {
	// Recovery middleware
	e.Use(
		gin.Recovery(),
		CROS(),
		gzip.Gzip(gzip.DefaultCompression),
	)

	// Logger middleware
	if config.App.Debug {
		pprof.Register(e)
		e.Use(gin.Logger())
	}

	return e
}
