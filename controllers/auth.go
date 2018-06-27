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

// Login 帳號和密碼
type Login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// NewAuthController constructor
func NewAuthController() *AuthController {
	return &AuthController{
		authService: services.NewAuthService(),
	}
}

// SignIn sign in
func (ctrl *AuthController) SignIn(c *gin.Context) {
	var login Login
	if err := c.ShouldBindJSON(&login); err == nil {
		if ctrl.authService.Attempt(login.Username, login.Password) {

			token, err := ctrl.authService.GenerateToken(login.Username)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "token generation failed",
					"token":   nil,
				})
			}

			c.JSON(http.StatusOK, gin.H{
				"message": "login successful",
				"token":   token,
			})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "login failed",
				"token":   nil,
			})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"token":   nil,
		})
	}
}
