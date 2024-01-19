package controller

import (
	"go-gin-api/src/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewStatus(server *gin.Engine, logger logger.Logrus) {

	server.GET("/status", func(ctx *gin.Context) {
		//middleware
		logger.Info("calling middleware status called")
	}, func(ctx *gin.Context) {

		ctx.JSON(http.StatusOK, gin.H{"status": "OK"})
	})

}
