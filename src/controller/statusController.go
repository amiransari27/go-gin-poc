package controller

import (
	"net/http"

	logger "github.com/openscriptsin/go-logger"

	"github.com/gin-gonic/gin"
)

// @Summary check status
// @Schemes
// @Description do ping
// @Tags Ping
// @Accept json
// @Produce json
// @Success 200 {string} NewStatus
// @Router /api/status [get]
func NewStatus(server *gin.Engine, logger logger.ILogrus) {

	api := server.Group("/api")
	api.GET("/status", func(ctx *gin.Context) {
		//middleware
		logger.Info(ctx, "calling middleware status called")
	}, func(ctx *gin.Context) {

		ctx.JSON(http.StatusOK, gin.H{"status": "OK"})
	})

}
