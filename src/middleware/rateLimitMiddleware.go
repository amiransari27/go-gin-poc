package middleware

import (
	"go-gin-api/src/config"
	"net/http"
	"sync"
	"time"

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

func GlobalIpBasedRateLimiterMiddleware() gin.HandlerFunc {
	rateLimiterConfig := config.GetConfig().RateLimiter

	if !rateLimiterConfig.Enabled {
		// if rate limiter is not enabled in cofig
		return func(ctx *gin.Context) { ctx.Next() }
	}

	type clientLimit struct {
		Limiter   *rate.Limiter
		LastVisit time.Time
	}

	var (
		clientsLimit = map[string]*clientLimit{}
		mu           sync.Mutex
		limiterRate  = rateLimiterConfig.Rate
		burst        = rateLimiterConfig.Burst
	)

	// background process, runs forevers
	// free up the limiter space after 1 mins of last visit
	go func() {
		for {
			time.Sleep(time.Minute)
			mu.Lock()
			for ip, limiter := range clientsLimit {
				if time.Since(limiter.LastVisit) > 3*time.Minute {
					delete(clientsLimit, ip)
				}
			}
			mu.Unlock()
		}
	}()

	return func(ctx *gin.Context) {
		// get IP of request
		clientId := ctx.ClientIP()
		mu.Lock()
		if _, found := clientsLimit[clientId]; !found {
			// creating rate limiter for each IP address
			clientsLimit[clientId] = &clientLimit{
				Limiter: rate.NewLimiter(limiterRate, burst),
			}
		}

		//setting last visit time
		clientsLimit[clientId].LastVisit = time.Now()

		if !clientsLimit[clientId].Limiter.Allow() {
			mu.Unlock()
			ctx.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"message": "Too many request"})
			return
		}

		mu.Unlock()
		ctx.Next()
	}

}
