package middleware

import (
	"fmt"
	"go-gin-api/src/service"
	"net/http"

	logger "github.com/openscriptsin/go-logger"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(jwtService service.IJWTService, logger logger.ILogrus) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		err := jwtService.TokenValid(ctx)

		if err != nil {
			fmt.Println("Invalid access")
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "Invalid access-token"})
			return
		}
		ctx.Next()
	}

}
