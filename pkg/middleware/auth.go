package middleware

import (
	"context"
	"go-boilerplate/pkg/auth"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5/request"
)

type UserService[T any] interface {
	Find(ctx context.Context, id string) (*T, error)
}

func AuthRequired[T any](jwt auth.JWTer, userSvc UserService[T]) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString, err := request.OAuth2Extractor.ExtractToken(ctx.Request)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, map[string]string{
				"error":             "invalid_request",
				"error_description": err.Error(),
			})
			return
		}

		token, err := jwt.ParseToken(tokenString)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, map[string]string{
				"error":             "invalid_request",
				"error_description": err.Error(),
			})
			return
		}

		subject, err := token.Claims.GetSubject()
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, map[string]string{
				"error":             "invalid_request",
				"error_description": err.Error(),
			})
			return
		}

		user, err := userSvc.Find(ctx, subject)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, map[string]string{
				"error":             "invalid_request",
				"error_description": err.Error(),
			})
			return
		}

		ctx.Set("user", user)
		ctx.Set("token", token)
		ctx.Next()
	}
}
