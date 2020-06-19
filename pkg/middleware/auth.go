package middleware

import (
	"go-boilerplate/pkg/auth"
	"go-boilerplate/pkg/config"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
)

// AuthRequired 認證JWT
func AuthRequired() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := request.ParseFromRequestWithClaims(
			ctx.Request,
			request.OAuth2Extractor,
			&jwt.StandardClaims{},
			func(token *jwt.Token) (interface{}, error) {
				return []byte(config.App.Key), nil
			},
		)

		if err == nil {
			authServ := auth.CreateService()
			u, err := authServ.GetUserFromToken(token.Raw)
			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error": err.Error(),
				})
				return
			}

			ctx.Set("user", u)
			ctx.Next()
			return
		}

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
}
