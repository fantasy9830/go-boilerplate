package routes

import (
	"go-boilerplate/pkg/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterWeb RegisterWeb
func (r *Router) RegisterWeb(web *gin.RouterGroup) error {
	web.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"name":    config.App.Name,
			"version": config.App.Version,
		})
	})

	return nil
}
