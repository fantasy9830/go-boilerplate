package middlewares

import (
	"net/http"

	"github.com/fantasy9830/go-boilerplate/Services"
	"github.com/fantasy9830/go-boilerplate/config"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
)

// Auth 認證JWT
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		conf := config.GetConfig()

		token, err := request.ParseFromRequestWithClaims(
			c.Request,
			request.AuthorizationHeaderExtractor,
			&services.Claims{},
			func(token *jwt.Token) (interface{}, error) {
				return []byte(conf.GetString("signingKey")), nil
			},
		)

		if err == nil {
			if token.Valid {
				auth := services.NewAuthService()
				claims, err := auth.ParseToken(token.Raw)

				if err != nil {
					c.JSON(http.StatusUnauthorized, gin.H{"status": "token is not valid"})
				} else {
					c.Set("claims", claims)

					c.Next()
				}
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"status": "token is not valid"})
			}
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized access to this resource"})
		}
	}
}
