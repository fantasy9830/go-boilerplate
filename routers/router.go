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
	// Logger middleware
	router.Use(gin.Logger())

	// Recovery middleware
	router.Use(gin.Recovery())

	// CORS middleware
	router.Use(middlewares.Cros())

	// 靜態目錄
	router.Static("/static", "./public")

	// favicon.ico
	router.StaticFile("/favicon.ico", "./public/favicon.ico")

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// 登入
	auth := controllers.NewAuthController()
	router.POST("/signin", auth.SignIn)

	// grpc
	grpc := &controllers.GrpcController{}
	router.GET("/grpc", grpc.SayHello)

	// 需認證
	authorized := router.Group("/")
	authorized.Use(middlewares.Auth())
	{
		authorized.GET("/test", func(c *gin.Context) {
			c.String(http.StatusOK, "Test")
		})
	}
}
