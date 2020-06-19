package auth

import (
	"github.com/gin-gonic/gin"
)

// Router Router
type Router func(*gin.RouterGroup)

// NewRoute NewRoute
func NewRoute(c *Controller) Router {
	return func(r *gin.RouterGroup) {
		r.POST("/oauth/token", c.OauthToken)
	}
}
