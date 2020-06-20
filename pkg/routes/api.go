package routes

import (
	"go-boilerplate/pkg/middleware"
	"go-boilerplate/pkg/websocket"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterAPI RegisterAPI
func (r *Router) RegisterAPI(api *gin.RouterGroup) error {
	v1 := api.Group("/v1")
	{
		v1.POST("/oauth/token", r.Auth.OauthToken)

		authorized := v1.Use(middleware.AuthRequired())
		{
			authorized.GET("/ws", func(ctx *gin.Context) {
				websocket.NewClient(ctx)
			})

			authorized.GET("/", func(ctx *gin.Context) {
				ctx.JSON(http.StatusOK, gin.H{
					"message": "API v1",
				})
			})
		}
	}

	return nil
}
