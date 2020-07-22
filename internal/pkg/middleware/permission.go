package middleware

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Error constants
var (
	ErrUnauthorized = errors.New("Unauthorized")
)

// Permission Permission
func Permission(permissions []string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, _, err := GetUserFromRequest(ctx.Request)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}

		for _, role := range user.Roles {
			for _, permission := range role.Permissions {
				for _, value := range permissions {
					if value == permission.Name {
						ctx.Next()
						return
					}
				}
			}
		}

		for _, permission := range user.Permissions {
			for _, value := range permissions {
				if value == permission.Name {
					ctx.Next()
					return
				}
			}
		}

		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": ErrUnauthorized.Error(),
		})
		return
	}
}

// Role Role
func Role(roles []string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, _, err := GetUserFromRequest(ctx.Request)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}

		for _, role := range user.Roles {
			for _, value := range roles {
				if value == role.Name {
					ctx.Next()
					return
				}
			}
		}

		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": ErrUnauthorized.Error(),
		})
		return
	}
}
