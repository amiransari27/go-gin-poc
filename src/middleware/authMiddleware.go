package middleware

import (
	"fmt"
	"go-gin-api/src/logger"
	"go-gin-api/src/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(ctx *gin.Context, jwtService service.IJWTService, logger logger.Logrus) {

	err := jwtService.TokenValid(ctx)

	if err != nil {
		fmt.Println("Invalid access")
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "Invalid access-token"})
		return
	}
	ctx.Next()

}
