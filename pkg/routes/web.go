package routes

import (
	"go-boilerplate/pkg/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func registerWeb(web *gin.RouterGroup) {
	web.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"name":    config.App.Name,
			"version": config.App.Version,
		})
	})
}
