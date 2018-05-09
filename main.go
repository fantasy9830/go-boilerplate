package main

import (
	// "github.com/fantasy9830/go-boilerplate/database/migrations"
	"github.com/fantasy9830/go-boilerplate/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	// setup mode
	gin.SetMode(gin.DebugMode)

	// get router instance
	router := routers.GetRouter()

	// setup router
	routers.SetupRouter()

	// run the migrations
	// migrations.Run()

	router.Run()
}
