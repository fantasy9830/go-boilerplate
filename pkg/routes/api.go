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
			user := v1.Group("user")
			user.GET("profile", r.User.Profile)

			// TODO: user list
			users := v1.Group("users")
			users.GET("", func(ctx *gin.Context) {
				data := make([]string, 0)
				ctx.JSON(http.StatusOK, data)
			})

			// TODO: permissions list
			permissions := v1.Group("permissions")
			permissions.GET("", func(ctx *gin.Context) {
				data := make([]string, 0)
				ctx.JSON(http.StatusOK, data)
			})

			// TODO: roles list
			roles := v1.Group("roles")
			roles.GET("", func(ctx *gin.Context) {
				data := make([]string, 0)
				ctx.JSON(http.StatusOK, data)
			})

			v1.GET("/ws", func(ctx *gin.Context) {
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
