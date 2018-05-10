package routers

import (
	"net/http"
	"sync"

	"github.com/fantasy9830/go-boilerplate/controllers"
	"github.com/fantasy9830/go-boilerplate/middlewares"
	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
	once   sync.Once
)

func init() {
	GetRouter()
}

// GetRouter gets the global router instance.
func GetRouter() *gin.Engine {
	once.Do(func() {
		router = gin.Default()
	})

	return router
}

// SetupRouter setup router
func SetupRouter() {

	// 認證的 Middleware
	router.Use(middlewares.Auth())

	// 靜態目錄
	router.Static("/static", "./public")

	// favicon.ico
	router.StaticFile("/favicon.ico", "./public/favicon.ico")

	// test
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// auth
	auth := &controllers.AuthController{}
	router.POST("/signin", auth.SignIn)

	// grpc
	grpc := &controllers.GrpcController{}
	router.GET("/grpc", grpc.SayHello)
}
