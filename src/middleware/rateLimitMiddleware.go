package middleware

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func GlobalRateLimiterMiddleware() gin.HandlerFunc {
	limiter := rate.NewLimiter(2, 4)
	return func(ctx *gin.Context) {
		if !limiter.Allow() {
			ctx.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"message": "Too many request"})
			return
		}
		ctx.Next()
	}

}

func IpBasedRateLimiterMiddleware() gin.HandlerFunc {
	var clientsIps = map[string]*rate.Limiter{}
	var mu sync.Mutex

	return func(ctx *gin.Context) {
		// get IP of request
		clientId := ctx.ClientIP()
		mu.Lock()
		if _, found := clientsIps[clientId]; !found {
			clientsIps[clientId] = rate.NewLimiter(2, 4)
		}

		if !clientsIps[clientId].Allow() {
			mu.Unlock()
			ctx.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"message": "Too many request"})
			return
		}

		mu.Unlock()
		ctx.Next()
	}

}
