package controllers

import (
	"errors"
	"fmt"
	"go-boilerplate/internal/app/models"
	"go-boilerplate/internal/app/services"
	"go-boilerplate/pkg/auth"
	"go-boilerplate/pkg/config"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Error constants
var (
	ErrSignatureInvalid = errors.New("Signature is invalid")
)

// Credentials Credentials
type Credentials struct {
	GrantType    string `form:"grant_type" json:"grant_type" binding:"required"`
	Username     string `form:"username" json:"username"`
	Password     string `form:"password" json:"password"`
	RefreshToken string `form:"refresh_token" json:"refresh_token"`
}

// Register Register
type Register struct {
	Name                 string `form:"name" json:"name" binding:"required"`
	Username             string `form:"username" json:"username" binding:"required"`
	Email                string `form:"email" json:"email" binding:"email" binding:"required"`
	Password             string `form:"password" json:"password" binding:"required,eqfield=PasswordConfirmation"`
	PasswordConfirmation string `form:"password_confirmation" json:"password_confirmation" binding:"required"`
}

// OAuthToken OAuthToken
type OAuthToken struct {
	AccessToken  string        `json:"access_token"`
	TokenType    string        `json:"token_type,omitempty"`
	RefreshToken string        `json:"refresh_token,omitempty"`
	ExpiresIn    time.Duration `json:"expires_in,omitempty"`
}

// AuthController Auth Controller
type AuthController struct {
	serv *services.AuthService
}

// NewAuthController New Auth Controller
func NewAuthController(service *services.AuthService) *AuthController {
	return &AuthController{
		serv: service,
	}
}

// OauthToken OauthToken
func (c *AuthController) OauthToken(ctx *gin.Context) {
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

		token, _, err := auth.CreateToken(user.ID, config.App.Key)
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
		refreshToken, err := auth.RefreshToken(credentials.RefreshToken, config.App.Key)
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

// Register Register
func (c *AuthController) Register(ctx *gin.Context) {
	var register Register
	if err := ctx.ShouldBind(&register); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := c.serv.Register(models.User{
		Name:     register.Name,
		Username: register.Username,
		Email:    register.Email,
		Password: register.Password,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
	return
}

// EmailVerify EmailVerify
func (c *AuthController) EmailVerify(ctx *gin.Context) {
	var dataURI struct {
		ID uint `uri:"id" binding:"required"`
	}

	var dataQuery struct {
		Signature string `form:"signature" binding:"required"`
	}

	if err := ctx.ShouldBindUri(&dataURI); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctx.ShouldBindQuery(&dataQuery); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := auth.ParseToken(dataQuery.Signature, config.App.Key)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	user, err := c.serv.EmailVerify(dataURI.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user.Password = ""

	ctx.JSON(http.StatusOK, user)
	return
}

// PasswordEmail PasswordEmail
func (c *AuthController) PasswordEmail(ctx *gin.Context) {
	var dataForm struct {
		Username string `form:"username" json:"username"`
	}

	if err := ctx.ShouldBind(&dataForm); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	url, err := c.serv.SendResetLink(dataForm.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, url)
	return
}

// PasswordReset PasswordReset
func (c *AuthController) PasswordReset(ctx *gin.Context) {
	var dataForm struct {
		Token                string `form:"token" json:"token" binding:"required"`
		Password             string `form:"password" json:"password" binding:"required,eqfield=PasswordConfirmation"`
		PasswordConfirmation string `form:"password_confirmation" json:"password_confirmation" binding:"required"`
	}

	if err := ctx.ShouldBind(&dataForm); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, err := c.serv.PasswordReset(dataForm.Token, dataForm.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, data)
	return
}
