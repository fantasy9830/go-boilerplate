package middleware

import (
	"net/http"
	"time"

	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"github.com/gin-gonic/gin"
)

var rateLimiter = tollbooth.NewLimiter(1, &limiter.ExpirableOptions{DefaultExpirationTTL: time.Hour})

// RateLimit RateLimit
func RateLimit() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := tollbooth.LimitByRequest(rateLimiter, ctx.Writer, ctx.Request); err != nil {
			ctx.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": err.Message,
			})
			return
		}

		ctx.Next()
		return
	}
}
