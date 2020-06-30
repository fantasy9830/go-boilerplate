package routes

import (
	"go-boilerplate/internal/pkg/middleware"
	"go-boilerplate/pkg/websocket"

	"github.com/gin-gonic/gin"
)

// RegisterWebSocket RegisterWebSocket
func (r *Router) RegisterWebSocket(ws *gin.RouterGroup) error {
	ws.Use(middleware.AuthRequired())
	{
		ws.GET("", func(ctx *gin.Context) {
			websocket.NewClient(ctx)
		})
	}

	return nil
}
