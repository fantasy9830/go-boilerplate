package controllers

import (
	"net/http"

	services "github.com/fantasy9830/go-boilerplate/Services"
	"github.com/gin-gonic/gin"
)

// AuthController ...
type AuthController struct {
	authService *services.AuthService
}

// NewAuthController constructor
func NewAuthController() *AuthController {
	return &AuthController{
		authService: services.NewAuthService(),
	}
}

// SignIn sign in
func (ctrl *AuthController) SignIn(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	if ctrl.authService.Attempt(username, password) {

		token, err := ctrl.authService.GenerateToken(username)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status": "token generation failed",
				"token":  nil,
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "you are logged in",
			"token":  token,
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": "unauthorized",
			"token":  nil,
		})
	}
}
