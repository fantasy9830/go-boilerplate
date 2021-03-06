package routes

import (
	"go-boilerplate/internal/pkg/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterAPI RegisterAPI
func (r *Router) RegisterAPI(api *gin.RouterGroup) error {
	v1 := api.Group("/v1")
	{
		v1.Use(middleware.RateLimit())
		{
			v1.POST("/oauth/token", r.Auth.OauthToken)
			v1.POST("/register", r.Auth.Register)
			v1.POST("/email/resend", r.Auth.EmailResend)
			v1.POST("/email/verify/:id", r.Auth.EmailVerify)
			v1.POST("/password/email", r.Auth.PasswordEmail)
			v1.POST("/password/reset", r.Auth.PasswordReset)

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

				authorized.GET("", func(ctx *gin.Context) {
					ctx.JSON(http.StatusOK, gin.H{
						"message": "API v1",
					})
				})
			}
		}
	}

	return nil
}
