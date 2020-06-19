package routes

import (
	"go-boilerplate/pkg/auth"
	"go-boilerplate/pkg/middleware"
	"go-boilerplate/pkg/websocket"
	"net/http"

	"github.com/gin-gonic/gin"
)

func registerAPI(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		auth.InitRoute()(v1)

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
}
