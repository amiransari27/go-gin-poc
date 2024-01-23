package middleware

import (
	"go-gin-api/src/appConst"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ClsMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		newId := uuid.New().String()
		requestId := ctx.Request.Header.Get(appConst.XRequestId)
		if requestId == "" {
			ctx.Set(appConst.XRequestId, newId)
		}
		ctx.Next()
	}
}
