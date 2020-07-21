package middleware

import (
	"go-boilerplate/internal/app/models"
	"go-boilerplate/internal/app/services"
	"go-boilerplate/pkg/config"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
)

// AuthRequired 認證JWT
func AuthRequired() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, _, err := GetUserFromRequest(ctx.Request)
		if err != nil {
			// var validationError *jwt.ValidationError
			// if errors.As(err, &validationError) {
			// 	claims := token.Claims.(*jwt.StandardClaims)

			// 	refreshExpire := time.Unix(claims.IssuedAt, 0).Add((config.App.RefreshTTL * time.Second))
			// 	if validationError.Errors == jwt.ValidationErrorExpired && time.Now().Before(refreshExpire) {
			// 		u, err := user.GetUserFromToken(token.Raw)
			// 		if err != nil {
			// 			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			// 				"error": err.Error(),
			// 			})
			// 			return
			// 		}

			// 		ctx.Set("user", u)
			// 		refreshToken, err := auth.RefreshToken(token.Raw)
			// 		if err != nil {
			// 			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			// 				"error": err.Error(),
			// 			})
			// 			return
			// 		}

			// 		// TODO: Blacklisting old token

			// 		ctx.Writer.Header().Add("Authorization", fmt.Sprintf("Bearer %s", refreshToken))
			// 		ctx.Next()
			// 		return
			// 	}
			// }

			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.Set("user", user)
		ctx.Next()
		return
	}
}

// GetUserFromRequest GetUserFromRequest
func GetUserFromRequest(req *http.Request) (user *models.User, token *jwt.Token, err error) {
	token, err = request.ParseFromRequestWithClaims(
		req,
		request.OAuth2Extractor,
		&jwt.StandardClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(config.App.Key), nil
		},
	)
	if err != nil {
		return
	}

	authServ := services.CreateAuthService()
	user, err = authServ.GetUserFromToken(token.Raw)

	return
}
