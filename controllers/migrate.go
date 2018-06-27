package controllers

import (
	"net/http"
	"time"

	"github.com/fantasy9830/go-boilerplate/database/migrations"
	"github.com/gin-gonic/gin"
)

// MigrateController ...
type MigrateController struct{}

// Run run the migrations
func (ctrl *MigrateController) Run(c *gin.Context) {
	migrations.Run()

	c.String(http.StatusOK, "complete "+time.Now().Format("2006-01-02 15:04:05"))
}

// Reverse reverse the migrations
func (ctrl *MigrateController) Reverse(c *gin.Context) {
	migrations.Reverse()

	c.String(http.StatusOK, "complete "+time.Now().Format("2006-01-02 15:04:05"))
}
