package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthController ...
type AuthController struct{}

// SignIn ...
func (ctrl *AuthController) SignIn(c *gin.Context) {
	data := "Sign In"

	c.JSON(http.StatusOK, gin.H{"data": data})
}
