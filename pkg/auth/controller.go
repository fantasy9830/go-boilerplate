package auth

import (
	"fmt"
	"go-boilerplate/pkg/config"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Credentials Credentials
type Credentials struct {
	GrantType    string `form:"grant_type" json:"grant_type" binding:"required"`
	Username     string `form:"username" json:"username"`
	Password     string `form:"password" json:"password"`
	RefreshToken string `form:"refresh_token" json:"refresh_token"`
}

// OAuthToken OAuthToken
type OAuthToken struct {
	AccessToken  string        `json:"access_token"`
	TokenType    string        `json:"token_type,omitempty"`
	RefreshToken string        `json:"refresh_token,omitempty"`
	ExpiresIn    time.Duration `json:"expires_in,omitempty"`
}

// Controller Controller
type Controller struct {
	serv *Service
}

// NewController New Controller
func NewController(service *Service) *Controller {
	return &Controller{
		serv: service,
	}
}

// OauthToken OauthToken
func (c *Controller) OauthToken(ctx *gin.Context) {
	var credentials Credentials
	if err := ctx.ShouldBind(&credentials); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if credentials.GrantType == "password" {
		user, err := c.serv.Attempt(credentials.Username, credentials.Password)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		token, _, err := CreateToken(user.ID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.Writer.Header().Set("Pragma", "no-cache")
		ctx.Writer.Header().Set("Cache-Control", "no-store")
		ctx.JSON(http.StatusOK, OAuthToken{
			AccessToken: token,
			TokenType:   "Bearer",
			ExpiresIn:   config.App.TTL,
		})
		return
	} else if credentials.GrantType == "refresh_token" {
		refreshToken, err := RefreshToken(credentials.RefreshToken)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.Writer.Header().Set("Pragma", "no-cache")
		ctx.Writer.Header().Set("Cache-Control", "no-store")
		ctx.JSON(http.StatusOK, OAuthToken{
			AccessToken:  refreshToken,
			TokenType:    "Bearer",
			ExpiresIn:    config.App.TTL,
			RefreshToken: credentials.RefreshToken,
		})

		return
	}

	ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Unsupported grant type: '%s'", credentials.GrantType)})
	return
}
