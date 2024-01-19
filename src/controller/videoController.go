package controller

import (
	"go-gin-api/src/entity"
	"go-gin-api/src/logger"
	"go-gin-api/src/middleware"
	"go-gin-api/src/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type VideoController interface {
	Save(ctx *gin.Context) error
	FindAll() []entity.Video
}

type videoController struct {
	service service.VideoService
	logger  logger.Logrus
}

func NewVideo(server *gin.Engine, service service.VideoService, jwtService service.IJWTService, logger logger.Logrus) {

	videoController := &videoController{
		service,
		logger,
	}

	group := server.Group("/videos", func(ctx *gin.Context) {
		middleware.AuthMiddleware(ctx, jwtService, logger)
	})

	group.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, videoController.findAll())
	})

	group.POST("/", func(ctx *gin.Context) {
		err := videoController.save(ctx)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "Video Input is valid!!"})
		}

	})

}

func (c *videoController) save(ctx *gin.Context) error {
	c.logger.Info("called save video")
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return err
	}

	r := c.service.Save(video)
	c.logger.Debug(r)
	return nil
}
func (c *videoController) findAll() []entity.Video {
	c.logger.Info("called find all videos")
	return c.service.FindAll()
}
