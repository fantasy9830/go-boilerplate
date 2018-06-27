package controllers

import (
	"net/http"
	"time"

	"github.com/fantasy9830/go-boilerplate/database/seeds"
	"github.com/gin-gonic/gin"
)

// SeederController ...
type SeederController struct{}

// Run run the seeds.
func (ctrl *SeederController) Run(c *gin.Context) {
	seeds.Run()

	c.String(http.StatusOK, "complete "+time.Now().Format("2006-01-02 15:04:05"))
}
